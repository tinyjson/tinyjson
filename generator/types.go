package generator

import (
	"fmt"
	"go/ast"
	"io/ioutil"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	typeDictionary = map[string]*TypeMeta{}
}

var (
	// typeDictionary is map from package absolute path & type name to meta.
	typeDictionary map[string]*TypeMeta
)

func SetMeta(path string, meta *TypeMeta) {
	typeDictionary[path+"#"+meta.name()] = meta
}

func Meta(path, name string) *TypeMeta {

	if baseType(name) {
		path = ""
	}
	_ = ParsePackage(path)

	meta := typeDictionary[path+"#"+name]
	return meta
}

func MetaRel(filePath, pkgName, typeName string) *TypeMeta {
	var pkgPath string
	if pkgName == "" {
		pkgPath = path.Dir(filePath)
	} else {
		pkgPath = resolvePackagePath(filePath, pkgName)
	}
	return Meta(pkgPath, typeName)
}

func baseType(name string) bool {
	return name == "string" || name == "bool" || name == "float" || name == "float64" ||
		name == "int" || name == "int8" || name == "int16" || name == "int32" || name == "int64" ||
		name == "uint" || name == "uint8" || name == "uint16" || name == "uint32" || name == "uint64"
}

type TypeMeta struct {
	//expr     ast.Expr
	spec *ast.TypeSpec

	// filePath is the place where type is defined. It is empty for base types
	// and slices.
	filePath string
	// pkgName is the package name from filePath.
	pkgName string

	pkgPath string
	imports map[string]string

	links []*TypeMeta

	// marshalFuncName, unmarshalFuncName names for function that marshal or unmarshal.
	marshalFuncName, unmarshalFuncName string
	marshalFunc                        *jen.Statement
	marshalFuncExternal                *jen.Statement
	unmarshalFunc                      *jen.Statement
	unmarshalFuncExternal              *jen.Statement

	//marshalCode func(assignVariable *jen.Statement) *jen.Statement

	// publicMarshal defines either create public marshal method either not.
	publicMarshal bool
	// publicUnmarshal defines either create public unmarshal method either not.
	publicUnmarshal bool

	// either can create _tinyjson.go file either not
	writable bool
}

func (meta *TypeMeta) name() string {
	return meta.spec.Name.Name
}

func (meta *TypeMeta) pkg() string {
	abs := meta.pkgPath
	var pkg string
	for {
		base := path.Base(pkg)
		if base == "vendor" || abs == path.Join(gopath, "src") || abs == path.Join(goroot, "src") {
			break
		}
		abs = path.Dir(abs)
		pkg = path.Join(base, pkg)
	}
	return pkg
}

func (meta *TypeMeta) nameType() ast.Expr {
	return meta.spec.Type
}

func Func(name, pkgPath string) *jen.Statement {
	return nil
}

func pointerLevel(expr ast.Expr) (ast.Expr, int) {
	var lvl int
	for {
		v, ok := expr.(*ast.StarExpr)
		if !ok {
			break
		}
		expr = v.X
		lvl++
	}
	return expr, lvl
}

func (meta *TypeMeta) buildMarshalFunc() {
	if meta.marshalFunc != nil {
		return
	}
	marsh := meta.buildMarshalCode(meta.spec.Type, jen.Id("this"), true, false)
	if marsh != nil {
		meta.marshalFunc = jen.Func().Id(meta.marshalFuncName).Call(
			jen.Id("w").Op("*").Qual("bytes", "Buffer"),
			jen.Id("this").Op("*").Add(jen.Id(meta.name())),
		).Block(marsh)
		meta.marshalFuncExternal = jen.Func().Id(meta.marshalFuncName).Call(
			jen.Id("w").Op("*").Qual("bytes", "Buffer"),
			jen.Id("this").Op("*").Add(jen.Qual(pkg(meta.filePath), meta.name())),
		).Block(marsh)
	}
}

func (meta *TypeMeta) buildUnmarshalFunc() {
	if meta.unmarshalFunc != nil {
		return
	}
	marsh, _ := meta.buildUnmarshalCode(meta.spec.Type, jen.Id("this"), jen.Id(meta.name()), true, false)
	if marsh != nil {
		meta.unmarshalFunc = jen.Func().Id(meta.unmarshalFuncName).Call(
			jen.Id("lex").Op("*").Qual("github.com/tinyjson/lexer", "Lexer"),
			jen.Id("this").Op("*").Add(jen.Id(meta.name())),
		).Error().Block(marsh, jen.Return(jen.Nil()))
		meta.unmarshalFuncExternal = jen.Func().Id(meta.unmarshalFuncName).Call(
			jen.Id("lex").Op("*").Qual("github.com/tinyjson/lexer", "Lexer"),
			jen.Id("this").Op("*").Add(jen.Qual(pkg(meta.filePath), meta.name())),
		).Error().Block(marsh, jen.Return(jen.Nil()))
	}
}

func (meta *TypeMeta) BuildPublicMarshalFunc() *jen.Statement {
	return jen.Func().Params(jen.Id("this").Op("*").Id(meta.name())).Id("MarshalJSON").
		Params().Params(jen.Id("[]byte"), jen.Error()).Block(
		jen.Id("w").Op(":=").Qual("bytes", "NewBuffer").Call(jen.Nil()),
		jen.Id(meta.marshalFuncName).Call(jen.Id("w"), jen.Id("this")),
		jen.Return(jen.Id("w").Dot("Bytes").Call(), jen.Nil()),
	)
}

func (meta *TypeMeta) BuildPublicUnmarshalFunc() *jen.Statement {
	return jen.Func().Params(jen.Id("this").Id("*"+meta.name())).Id("UnmarshalJSON").
		Params(jen.Id("data").Id("[]byte")).Params(jen.Error()).Block(
		jen.Id("lex").Op(":=").Qual("github.com/tinyjson/lexer", "NewLexer").Call(jen.Id("data")),
		jen.Id("lex").Dot("Parse").Call(),
		jen.Return(jen.Id(meta.unmarshalFuncName).Call(jen.Id("lex"), jen.Id("this"))),
	)
}

func (meta *TypeMeta) AddLink(link *TypeMeta) {
	for _, v := range meta.links {
		if v == link {
			return
		}
	}
	meta.links = append(meta.links, link)
}

func (meta *TypeMeta) buildMarshalCode(expr ast.Expr, marshalingVariable *jen.Statement, pointer bool, inFor bool) (code *jen.Statement) {
	var lvl int
	expr, lvl = pointerLevel(expr)
	if lvl > 0 {
		fmt.Printf("It is not possible to create marshal method for pointer types. Type %s in %s ignored for export.\n", meta.name(), meta.filePath)
		if meta.publicMarshal {
			meta.publicMarshal = false
		}
	}
	code = jen.Empty()
	for i := lvl; i > 0; i-- {
		code = code.If(jen.Op("*").Add(marshalingVariable).Op("==").Nil()).Block(
			jen.Id("w").Dot("WriteString").Call(jen.Id(`"null"`)),
			jen.Return(),
		).Line()
		v := Variable()
		code = code.Add(v).Op(":=").Op("*").Add(marshalingVariable).Line()
		marshalingVariable = v
	}

	marshalingValueVariable := marshalingVariable
	marshalingValuePointer := marshalingVariable
	if pointer {
		marshalingValueVariable = jen.Op("*").Add(marshalingValueVariable)
	} else {
		marshalingValuePointer = jen.Op("&").Add(marshalingValuePointer)
	}

	switch v := expr.(type) {
	case *ast.Ident:
		switch v.Name {
		case "string":
			code = code.Id("w").Dot("WriteString").Call(
				jen.Qual("strconv", "Quote").Call(jen.String().Call(marshalingValueVariable)))
		case "int", "int8", "int16", "int32", "int64":
			code = code.Id("w").Dot("WriteString").Call(
				jen.Qual("strconv", "FormatInt").
					Call(jen.Int64().Call(marshalingValueVariable), jen.Id("10")))
		case "uint", "uint8", "uint16", "uint32", "uint64":
			code = code.Id("w").Dot("WriteString").Call(
				jen.Qual("strconv", "FormatUint").
					Call(jen.Uint64().Call(marshalingValueVariable), jen.Id("10")))
		case "float32", "float64":
			code = code.Id("w").Dot("WriteString").Call(
				jen.Qual("strconv", "FormatFloat").
					Call(jen.Float64().Call(marshalingValueVariable), jen.Id("'g'"), jen.Id("-1"), jen.Id("64")))
		case "bool":
			code = code.Id("w").Dot("WriteString").Call(
				jen.Qual("strconv", "FormatBool").Call(jen.Bool().Call(marshalingValueVariable)))
		default: // alias
			aliasMeta := MetaRel(meta.filePath, "", v.Name)
			if aliasMeta == nil {
				return nil
			}
			// TODO if nil - type not found
			meta.AddLink(aliasMeta)
			code = code.Id(aliasMeta.marshalFuncName).Call(jen.Id("w"), jen.Call(jen.Op("*").Id(aliasMeta.name())).Call(marshalingValuePointer))
			//code = code.Id(aliasMeta.marshalFuncName).Call(jen.Id("w"), jen.Op("&").Add(marshalingVariable))
		}
		return
	case *ast.SelectorExpr: // external alias
		//x, lvl := pointerLevel(v.X)

		x, ok := v.X.(*ast.Ident)
		if !ok {
			fmt.Printf("unknown type: %#v\n", x)
			return nil
		}
		pkg := meta.imports[x.Name]

		aliasMeta := MetaRel(meta.filePath, pkg, v.Sel.Name)
		if aliasMeta == nil {
			fmt.Println("alias not found: ", meta.pkgPath, pkg, v.Sel.Name)
			return
		}
		if aliasMeta.publicMarshal {
			// TODO: test
			code = code.List(jen.Id("data"), jen.Id("_")).Op("=").
				Add(marshalingVariable).Call(jen.Qual(pkg, v.Sel.Name)).Dot("MarshalJSON").Call().
				Line().
				Id("w").Dot("Write").Call(jen.Id("data"))
		} else {
			meta.AddLink(aliasMeta)
			code = code.Id(aliasMeta.marshalFuncName).Call(jen.Id("w"), jen.Call(jen.Op("*").Qual(pkg, aliasMeta.name())).Call(marshalingVariable))
		}
		return
	case *ast.ArrayType:
		var breakSt *jen.Statement
		if inFor {
			breakSt = jen.Break()
		} else {
			breakSt = jen.Return()
		}
		code = code.If(marshalingValueVariable.Clone().Op("==").Nil()).Block(
			jen.Id("w").Dot("WriteString").Call(jen.Id(`"null"`)),
			breakSt,
		).Line()

		//fmt.Printf("array: %#v %#v\n", v.Elt, v.Len)
		varI := Variable()
		varV := Variable()

		code = code.Id("w").Dot("WriteString").Call(jen.Id(`"["`)).Line()
		code = code.For(jen.List(varI, varV).Op(":=").Range().Add(marshalingValueVariable)).Block(
			jen.If(varI.Clone().Op(">").Id("0")).Block(
				jen.Id("w").Dot("WriteString").Call(jen.Id(`","`)),
			),
			meta.buildMarshalCode(v.Elt, varV, false, true),
		).Line()
		code = code.Id("w").Dot("WriteString").Call(jen.Id(`"]"`))
		return code
	case *ast.MapType:
		var breakSt *jen.Statement
		if inFor {
			breakSt = jen.Break()
		} else {
			breakSt = jen.Return()
		}
		code = code.If(marshalingValueVariable.Clone().Op("==").Nil()).Block(
			jen.Id("w").Dot("WriteString").Call(jen.Id(`"null"`)),
			breakSt,
		).Line()

		varKey := Variable()
		varVal := Variable()

		varNotFirst := Variable()

		code = code.Var().Add(varNotFirst).Bool().Line()
		code = code.Id("w").Dot("WriteString").Call(jen.Id(`"{"`)).Line()
		code = code.For(jen.List(varKey, varVal)).Op(":=").Range().Add(marshalingValueVariable).Block(
			jen.If(varNotFirst).Block(
				jen.Id("w").Dot("WriteString").Call(jen.Id(`","`)),
			).Else().Block(
				varNotFirst.Clone().Op("=").True(),
			),
			buildMapKey(varKey, v.Key),
			jen.Id("w").Dot("WriteString").Call(jen.Id(`":"`)),
			meta.buildMarshalCode(v.Value, varVal, false, true),
		).Line()
		code = code.Id("w").Dot("WriteString").Call(jen.Id(`"}"`))
		return code
	case *ast.StructType:
		code = code.Id("w").Dot("WriteString").Call(jen.Id(`"{"`))

		mp := newStructKeyMapper(v.Fields, marshalingVariable)
		mp.Proc(meta)

		for i, name := range mp.Keys() {
			field := mp.dic[name]
			if i > 0 {
				code = code.Line().Id("w").Dot("WriteString").Call(jen.Id(`","`))
			}
			code = code.Line().Id("w").Dot("WriteString").Call(jen.Id(strconv.Quote(strconv.Quote(name) + ":")))
			code = code.Line().Add(meta.buildMarshalCode(field.typeExpr, field.variableName, false, true))
		}

		code = code.Line().Id("w").Dot("WriteString").Call(jen.Id(`"}"`))
		return code
	default:
		fmt.Printf("Unsupported type %s in %s, please report.\n", meta.name(), meta.filePath)
		if meta.publicMarshal {
			meta.publicMarshal = false
		}
	}

	return nil
}

func (meta *TypeMeta) buildUnmarshalCode(expr ast.Expr, marshalingVariable, marshalingVariableType *jen.Statement, pointer bool, inFor bool) (code, value *jen.Statement) {
	var lvl int
	variableType, valueLvl := buildTypeDeclaration(expr, meta.imports)
	expr, lvl = pointerLevel(expr)
	if lvl > 0 {
		fmt.Printf("It is not possible to create unmarshal method for pointer types. Type %s in %s ignored for Unmarshal export.\n", meta.name(), meta.filePath)
		if meta.publicUnmarshal {
			meta.publicUnmarshal = false
		}
	}
	code = jen.Empty()

	var (
		unmarshalingValueVariable        *jen.Statement
		unmarshalingValuePointer         *jen.Statement
		marshalingVariableInitialization bool
	)
	marshalingVariableInitiated := Variable()
	if marshalingVariable == nil {
		if pointer {
			unmarshalingValueVariable = jen.Op("*").Add(marshalingVariableInitiated)
			unmarshalingValuePointer = marshalingVariableInitiated
		} else {
			unmarshalingValueVariable = marshalingVariableInitiated
			unmarshalingValuePointer = jen.Op("&").Add(marshalingVariableInitiated)
		}
		defer func() {
			if marshalingVariableInitialization {
				if _, ok := expr.(*ast.ArrayType); ok {
					if pointer {
						code = jen.Var().Add(value).Op("*").Add(variableType).Line().Add(code)
					} else {
						code = jen.Var().Add(value).Add(variableType).Line().Add(code)
					}
				} else {
					if pointer {
						code = jen.Add(value).Op(":=").Op("&").Add(variableType).Op("{").Op("}").Line().Add(code)
					} else {
						code = jen.Add(value).Op(":=").Add(variableType).Op("{").Op("}").Line().Add(code)
					}
				}
			}
		}()
	} else {
		unmarshalingValueVariable = marshalingVariable
		unmarshalingValuePointer = marshalingVariable
		if pointer {
			unmarshalingValueVariable = jen.Op("*").Add(unmarshalingValueVariable)
		} else {
			unmarshalingValuePointer = jen.Op("&").Add(unmarshalingValuePointer)
		}
	}

	switch v := expr.(type) {
	case *ast.Ident:
		switch v.Name {
		case "string":
			varString := Variable()
			code.List(varString, jen.Err()).Op(":=").Id("lex").Dot("ReadString").Call().List().Line()
			value = marshalingVariableType.Clone().Call(jen.Add(varString))
			code.If(jen.Err().Op("!=").Nil()).Block(jen.Return(jen.Err())).Line()
			if marshalingVariable != nil {
				code.Add(unmarshalingValueVariable.Clone().Op("=").Call(marshalingVariableType).Call(varString)).Line()
			}
		case "int", "int8", "int16", "int32", "int64":
			varString := Variable()
			code.List(varString, jen.Err()).Op(":=").Id("lex").Dot("ReadInt").Call().List().Line()
			if valueLvl > 0 {
				code.If(jen.Err().Op("==").Qual("github.com/tinyjson/lexer", "ErrorNilValue")).Block(
					jen.Err().Op("=").Nil(),
					unmarshalingValueVariable.Clone().Op("=").Nil(),
					jen.Return(jen.Nil()),
				).Line()
				newVar := Variable()
				code.Add(newVar).Op(":=").Add(variableType).Call(varString).Line()
				varString = newVar
				for i := 0; i < valueLvl; i++ {
					newVar := Variable()
					code.Add(newVar).Op(":=").Op("&").Add(varString).Line()
					varString = newVar
				}
			}
			code.Add(unmarshalingValueVariable).Op("=").Add(marshalingVariableType).Call(jen.Add(varString)).Line()
			code.If(jen.Err().Op("!=").Nil()).Block(jen.Return(jen.Err()))
		case "uint", "uint8", "uint16", "uint32", "uint64":
			varString := Variable()
			code = jen.List(varString, jen.Err()).Op(":=").Id("lex").Dot("ReadInt").Call().List().Line()
			code.Add(unmarshalingValueVariable).Op("=").Add(marshalingVariableType).Call(jen.Add(varString)).Line()
			code.If(jen.Err().Op("!=").Nil()).Block(jen.Return(jen.Err()))
		case "float32", "float64":
			varString := Variable()
			code = jen.List(varString, jen.Err()).Op(":=").Id("lex").Dot("ReadFloat").Call().List().Line()
			code.Add(unmarshalingValueVariable).Op("=").Add(marshalingVariableType).Call(jen.Add(varString)).Line()
			code.If(jen.Err().Op("!=").Nil()).Block(jen.Return(jen.Err()))
		case "bool":
			varString := Variable()
			code = jen.List(varString, jen.Err()).Op(":=").Id("lex").Dot("ReadBool").Call().List().Line()
			code.Add(unmarshalingValueVariable).Op("=").Add(marshalingVariableType).Call(jen.Add(varString)).Line()
			code.If(jen.Err().Op("!=").Nil()).Block(jen.Return(jen.Err()))
		default: // alias
			aliasMeta := MetaRel(meta.filePath, "", v.Name)
			if aliasMeta == nil {
				return nil, nil
			}
			// TODO if nil - type not found
			meta.AddLink(aliasMeta)
			code = code.Id(aliasMeta.unmarshalFuncName).Call(jen.Id("lex"), jen.Call(jen.Op("*").Id(aliasMeta.name())).Call(unmarshalingValuePointer))
			//code = code.Id(aliasMeta.marshalFuncName).Call(jen.Id("w"), jen.Op("&").Add(marshalingVariable))
		}
		return
	case *ast.SelectorExpr: // external alias
		x, ok := v.X.(*ast.Ident)
		if !ok {
			fmt.Printf("unknown type: %#v\n", x)
			return nil, nil
		}
		pkg := meta.imports[x.Name]

		aliasMeta := MetaRel(meta.filePath, pkg, v.Sel.Name)
		if aliasMeta == nil {
			fmt.Println("alias not found: ", meta.pkgPath, pkg, v.Sel.Name)
			return
		}
		if aliasMeta.publicUnmarshal {
			// TODO: test
			//code = code.List(jen.Id("data"), jen.Id("_")).Op("=").
			//	Add(marshalingVariable).Call(jen.Qual(pkg, v.Sel.Name)).Dot("MarshalJSON").Call().
			//	Line().
			//	Id("w").Dot("Write").Call(jen.Id("data"))
		} else {
			meta.AddLink(aliasMeta)
			code = code.Id(aliasMeta.unmarshalFuncName).Call(jen.Id("lex"), jen.Call(jen.Op("*").Qual(pkg, aliasMeta.name())).Call(marshalingVariable))
		}
		return
	case *ast.ArrayType:
		marshalingVariableInitialization = true
		if marshalingVariable == nil {
			value = marshalingVariableInitiated
		}
		code.If(jen.Id("lex").Dot("Controls").Index(jen.Id("0")).Op("==").Qual("github.com/tinyjson/lexer", "Nil").Block(
			unmarshalingValueVariable.Clone().Op("=").Nil(),
			jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()),
		)).Else().If(jen.Id("lex").Dot("Controls").Index(jen.Id("0")).Op("!=").Qual("github.com/tinyjson/lexer", "ArrayIn").Block(
			jen.Id("lex").Dot("SkipValue").Call(),
			jen.Return(jen.Qual("github.com/tinyjson/lexer", "ErrorUnexpectedType")),
		))

		var arrayCode []jen.Code

		arrayCode = append(arrayCode)
		arrayCode = append(arrayCode, jen.Add(unmarshalingValueVariable).Op("=").Make(variableType, jen.Id("0")))

		arrayCode = append(arrayCode, jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()))
		arrayCode = append(arrayCode, jen.Id("lex").Dot("Actions").Op("=").Id("lex").Dot("Actions").Index(jen.Id("4"), jen.Empty()))

		t, _ := buildTypeDeclaration(v.Elt, meta.imports)
		ccode, variable := meta.buildUnmarshalCode(v.Elt, nil, t, false, true)

		arrayCode = append(arrayCode, jen.For().Block(
			jen.If(jen.Id("lex").Dot("Controls").Index(jen.Id("0")).Op("==").Qual("github.com/tinyjson/lexer", "ArrayOut")).Block(
				jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()).Line(),
				jen.Break(),
			),
			ccode,
			unmarshalingValueVariable.Clone().Op("=").Append(unmarshalingValueVariable, variable),
		))
		code.Else().Block(arrayCode...).Line()
		return
	case *ast.MapType:
		marshalingVariableInitialization = true
		if marshalingVariable == nil {
			value = marshalingVariableInitiated
		}
		code.Id("data").Op(":=").Id("lex").Dot("Data").Call().Line()
		code.If(jen.Id("lex").Dot("Controls").Index(jen.Id("0")).Op("==").Qual("github.com/tinyjson/lexer", "Nil").Block(
			unmarshalingValueVariable.Clone().Op("=").Nil(),
			jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()),
			//jen.Return(jen.Nil()),
		)).Else().If(jen.Id("lex").Dot("Controls").Index(jen.Id("0")).Op("!=").Qual("github.com/tinyjson/lexer", "ObjectIn").Block(
			jen.Id("lex").Dot("SkipValue").Call(),
			//jen.Return(jen.Qual("github.com/tinyjson/lexer", "ErrorUnexpectedType")),
		))

		var mapCode []jen.Code

		mapCode = append(mapCode, jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()))
		mapCode = append(mapCode, jen.Id("lex").Dot("Actions").Op("=").Id("lex").Dot("Actions").Index(jen.Id("4"), jen.Empty()))

		t, _ := buildTypeDeclaration(expr, meta.imports)
		mapCode = append(mapCode, jen.Add(unmarshalingValueVariable).Op("=").Make(t, jen.Id("0")))

		if pointer {
			unmarshalingValueVariable = jen.Call(unmarshalingValueVariable)
		}

		keyVariable := Variable()

		forBlock := []jen.Code{
			jen.If(jen.Id("lex").Dot("Controls").Index(jen.Id("0")).Op("==").Qual("github.com/tinyjson/lexer", "ObjectOut")).Block(
				jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()),
				jen.Break(),
			),
			jen.List(keyVariable, jen.Id("_")).Op(":=").Qual("strconv", "Unquote").Call(jen.Id(
				"string").Call(jen.Id("data").Index(jen.Id("lex").Dot("Actions").Index(jen.Id("0")), jen.Id("lex").Dot("Actions").Index(jen.Id("1"))))),
		}
		var ccode *jen.Statement
		ccode, keyVariable = parseMapKey(keyVariable, v.Key)
		if ccode != nil {
			forBlock = append(forBlock, ccode)
		}

		t, _ = buildTypeDeclaration(v.Value, meta.imports)
		ccode, variable := meta.buildUnmarshalCode(v.Value, nil, t, false, true)

		forBlock = append(forBlock,
			jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()),
			jen.Id("lex").Dot("Actions").Op("=").Id("lex").Dot("Actions").Index(jen.Id("2"), jen.Empty()),
			ccode,
			unmarshalingValueVariable.Clone().Index(keyVariable).Op("=").Add(variable),
		)

		mapCode = append(mapCode, jen.For().Block(forBlock...))
		code.Else().Block(mapCode...)
		return
	case *ast.StructType:
		code.Id("data").Op(":=").Id("lex").Dot("Data").Call().Line()
		code.If(jen.Id("lex").Dot("Controls").Index(jen.Id("0")).Op("==").Qual("github.com/tinyjson/lexer", "Nil").Block(
			jen.Return(jen.Nil()),
		)).Else().If(jen.Id("lex").Dot("Controls").Index(jen.Id("0")).Op("!=").Qual("github.com/tinyjson/lexer", "ObjectIn").Block(
			jen.Id("lex").Dot("SkipValue").Call(),
			jen.Return(jen.Qual("github.com/tinyjson/lexer", "ErrorUnexpectedType")),
		)).Line()

		code.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()).Line()
		code.Id("lex").Dot("Actions").Op("=").Id("lex").Dot("Actions").Index(jen.Id("4"), jen.Empty()).Line()

		mp := newStructKeyMapper(v.Fields, marshalingVariable)
		mp.Proc(meta)

		var switchKey []jen.Code
		switchKey = append(switchKey, jen.Case(jen.Qual("github.com/tinyjson/lexer", "ObjectOut")))
		switchKey = append(switchKey, jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()))
		switchKey = append(switchKey, jen.Return(jen.Nil()))
		switchKey = append(switchKey, jen.Case(jen.Qual("github.com/tinyjson/lexer", "Key")))
		switchKey = append(switchKey, jen.List(jen.Id("key"), jen.Id("_")).Op(":=").Qual("strconv", "Unquote").Call(jen.Id(
			"string").Call(jen.Id("data").Index(jen.Id("lex").Dot("Actions").Index(jen.Id("0")), jen.Id("lex").Dot("Actions").Index(jen.Id("1"))))))
		switchKey = append(switchKey, jen.Id("lex").Dot("Controls").Op("=").Id("lex").Dot("Controls").Index(jen.Id("1"), jen.Empty()))
		switchKey = append(switchKey, jen.Id("lex").Dot("Actions").Op("=").Id("lex").Dot("Actions").Index(jen.Id("2"), jen.Empty()))

		var switchValue []jen.Code

		for _, name := range mp.Keys() {
			field := mp.dic[name]
			switchValue = append(switchValue, jen.Case(jen.Id(strconv.Quote(name))))
			marshalingVariableType, _ := buildTypeDeclaration(field.typeExpr, meta.imports)
			ccode, _ := meta.buildUnmarshalCode(field.typeExpr, field.variableName, marshalingVariableType, false, true)
			switchValue = append(switchValue, ccode)
		}
		switchValue = append(switchValue, jen.Default())
		switchValue = append(switchValue, jen.Id("lex").Dot("SkipValue").Call())

		switchKey = append(switchKey, jen.Switch(jen.Id("key")).Block(switchValue...))

		code.For().Block(
			jen.Switch(jen.Id("lex").Dot("Controls").Index(jen.Id("0"))).Block(switchKey...),
		)
		return
	default:
		fmt.Printf("Unsupported type %s in %s, please report.\n", meta.name(), meta.filePath)
		if meta.publicUnmarshal {
			meta.publicUnmarshal = false
		}
	}

	return nil, nil
}

func WriteTypes() {
	fileTypes := map[string]map[*TypeMeta]struct{}{}
	for _, meta := range typeDictionary {
		types := fileTypes[meta.filePath+"#"+meta.pkgName]
		if types == nil {
			types = map[*TypeMeta]struct{}{}
		}
		if !meta.publicMarshal && !meta.publicUnmarshal {
			continue
		}

		origMeta := meta
		q := []*TypeMeta{meta}
		for {
			if len(q) == 0 {
				break
			}
			meta := q[0]
			q = q[1:]

			meta.buildMarshalFunc()
			meta.buildUnmarshalFunc()

			if origMeta != meta && origMeta.pkgPath == meta.pkgPath && origMeta.filePath != meta.filePath {
				// From the same package, but different file
				types := fileTypes[meta.filePath+"#"+origMeta.pkgName]
				if types == nil {
					types = map[*TypeMeta]struct{}{}
				}
				types[meta] = struct{}{}
				fileTypes[origMeta.filePath+"#"+meta.pkgName] = types
			} else {
				types[meta] = struct{}{}
			}
			q = append(q, meta.links...)
		}

		// TODO: fix if foo.go depends on an external lib and bar.go depends on the same external lib, then function will duplicate.
		fileTypes[origMeta.filePath+"#"+meta.pkgName] = types
	}

	for file, m := range fileTypes {
		parts := strings.Split(file, "#")
		file = parts[0]
		if len(m) == 0 {
			continue
		}
		f := jen.NewFile(parts[1])

		var metas []*TypeMeta

		for meta := range m {
			metas = append(metas, meta)
		}

		sort.Slice(metas, func(i, j int) bool {
			return metas[i].filePath+metas[i].name() < metas[j].filePath+metas[j].name()
		})

		for _, meta := range metas {
			if pkg(meta.filePath) == pkg(file) {
				if meta.marshalFunc != nil {
					f.Add(meta.marshalFunc).Line().Line()
				}
				if meta.unmarshalFunc != nil {
					f.Add(meta.unmarshalFunc).Line().Line()
				}
			} else {
				if meta.marshalFuncExternal != nil {
					f.Add(meta.marshalFuncExternal).Line().Line()
				}
				if meta.unmarshalFuncExternal != nil {
					f.Add(meta.unmarshalFuncExternal).Line().Line()
				}
			}

			if meta.publicMarshal {
				f.Add(meta.BuildPublicMarshalFunc().Line().Line())
			}
			if meta.publicUnmarshal {
				f.Add(meta.BuildPublicUnmarshalFunc().Line().Line())
			}
		}

		file = strings.TrimSuffix(file, ".go") + "_tinyjson.go"
		if err := ioutil.WriteFile(file, []byte(fmt.Sprintf("%#v", f)), 0644); err != nil {
			fmt.Printf("Could not write to file %v: %v.\n", file, err)
		}
	}
}
