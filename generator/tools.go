package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
)

func init() {
	fset = token.NewFileSet()
	usedFuncCodes = map[string]struct{}{}
	parsedPackages = map[string]struct{}{}
	gopath = os.Getenv("GOPATH")
	goroot = os.Getenv("GOROOT")
}

var variableCounter int64

func Variable() *jen.Statement {
	variableCounter++
	return jen.Id("v" + strconv.FormatInt(variableCounter, 10))
}

func buildMapKey(varKey *jen.Statement, keyType ast.Expr) *jen.Statement {
	var st *jen.Statement
	switch v := keyType.(type) {
	case *ast.Ident:
		switch v.Name {
		case "string":
			st = jen.Qual("strconv", "Quote").Call(jen.String().Call(varKey))
		case "int", "int8", "int16", "int32", "int64":
			st = jen.Id(`"\""`).Op("+").Qual("strconv", "FormatInt").
				Call(jen.Int64().Call(varKey), jen.Id("10")).Op("+").Id(`"\""`)
		case "uint", "uint8", "uint16", "uint32", "uint64":
			st = jen.Id(`"\""`).Op("+").Qual("strconv", "FormatUint").
				Call(jen.Uint64().Call(varKey), jen.Id("10")).Op("+").Id(`"\""`)
		case "float32", "float64":
			st = jen.Id(`"\""`).Op("+").Qual("strconv", "FormatFloat").
				Call(jen.Float64().Call(varKey), jen.Id("'g'"), jen.Id("-1"), jen.Id("64")).Op("+").Id(`"\""`)
		case "bool":
			st = jen.Id(`"\""`).Op("+").Qual("strconv", "FormatBool").Call(jen.Bool().Call(varKey)).Op("+").Id(`"\""`)
		default:
			// TODO: add support for type alias for base types
			// TODO: add support for MarshalText
			//"type alias"
			return nil
		}
		if st != nil {
			st = jen.Id("w").Dot("WriteString").Call(st)
		}
		return st

	}
	jen.Id("w").Dot("WriteString").Call(jen.Qual("strconv", "Quote").Call(varKey))
	return st
}

func parseMapKey(varKey *jen.Statement, keyType ast.Expr) (code *jen.Statement, key *jen.Statement) {
	switch v := keyType.(type) {
	case *ast.Ident:
		switch v.Name {
		case "string":
			key = varKey
		case "int", "int8", "int16", "int32", "int64":
			key = Variable()
			code = jen.List(key, jen.Id("_")).Op(":=").Qual("strconv", "ParseInt").
				Call(varKey, jen.Id("10"), jen.Id("64"))
			key = jen.Id(v.Name).Call(key)
		case "uint", "uint8", "uint16", "uint32", "uint64":
			key = Variable()
			code = jen.List(key, jen.Id("_")).Op(":=").Qual("strconv", "ParseUint").
				Call(varKey, jen.Id("10"), jen.Id("64"))
			key = jen.Id(v.Name).Call(key)
		case "float32", "float64":
			key = Variable()
			code = jen.List(key, jen.Id("_")).Op(":=").Qual("strconv", "ParseFloat").
				Call(varKey, jen.Id("64"))
			key = jen.Id(v.Name).Call(key)
		case "bool":
			key = Variable()
			code = jen.List(key, jen.Id("_")).Op(":=").Qual("strconv", "ParseBool").
				Call(varKey)
			key = jen.Id(v.Name).Call(key)
		default:
			// TODO: add support for type alias for base types
			// TODO: add support for MarshalText
			//"type alias"
			return nil, nil
		}

	}
	return
}

var goroot, gopath string

var exist = func(dirPath string) bool {
	_, err := os.Stat(dirPath)
	return !os.IsNotExist(err)
}

func resolvePackagePath(filePath, pkgName string) string {
	pkgPath := path.Dir(filePath)
	for {
		potential := path.Join(pkgPath, "vendor", pkgName)
		if exist(potential) {
			return potential
		}
		if pkgPath == "/" || pkgPath == "" || pkgPath == path.Join(gopath, "src") {
			break
		}
		pkgPath = path.Dir(pkgPath)
	}

	potential := path.Join(goroot, "src", pkgName)
	if exist(potential) {
		return potential
	}
	potential = path.Join(gopath, "src", pkgName)
	if exist(potential) {
		return potential
	}
	return ""
}

func pkg(filePath string) string {
	abs := path.Dir(filePath)
	var pkg string
	for {
		_, base := path.Split(abs)
		if base == "vendor" || abs == path.Join(gopath, "src") || abs == path.Join(goroot, "src") {
			break
		}
		abs = path.Dir(abs)
		pkg = path.Join(base, pkg)
	}
	return pkg
}

var (
	fset           *token.FileSet
	parsedPackages map[string]struct{}
)

func ParsePackage(pkgPath string) error {
	if _, ok := parsedPackages[pkgPath]; ok {
		return nil
	}
	parsedPackages[pkgPath] = struct{}{}

	pkg, err := parser.ParseDir(fset, pkgPath, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("could not parse dir %v: %v", pkgPath, err)
		return nil
	}

	var (
		writable bool
		metas    []*TypeMeta
	)
	for _, pkg := range pkg {
		for filePath, tree := range pkg.Files {
			imports := map[string]string{}
			for _, imp := range tree.Imports {
				pkgPath, _ := strconv.Unquote(imp.Path.Value)
				var name string
				if imp.Name != nil {
					name = imp.Name.Name
				} else {
					name = path.Base(pkgPath)
				}
				imports[name] = pkgPath
			}

			for _, obj := range tree.Scope.Objects {
				if obj.Kind != ast.Typ {
					continue
				}
				spec := obj.Decl.(*ast.TypeSpec)

				code := generateFuncCode()

				public := strings.Contains(spec.Doc.Text(), "tinyjson:json")

				meta := &TypeMeta{
					spec:              spec,
					filePath:          filePath,
					pkgName:           pkg.Name,
					pkgPath:           pkgPath,
					imports:           imports,
					marshalFuncName:   "tinyjsonMarshalC" + code,
					unmarshalFuncName: "tinyjsonUnmarshalC" + code,
					publicMarshal:     public,
					publicUnmarshal:   public,
				}
				SetMeta(pkgPath, meta)
				metas = append(metas, meta)
			}
		}
	}
	if writable {
		for _, meta := range metas {
			meta.writable = true
		}
	}

	return nil
}

var usedFuncCodes map[string]struct{}

func generateFuncCode() string {
	for {
		code := strconv.FormatInt(rand.Int63(), 16)
		if _, ok := usedFuncCodes[code]; ok {
			continue
		}
		usedFuncCodes[code] = struct{}{}

		return code
	}
}

func buildTypeDeclaration(expr ast.Expr, imports map[string]string) (*jen.Statement, int) {
	st := jen.Empty()
	var lvl int
	expr, lvl = pointerLevel(expr)
	//for i := 0; i < lvl; i++ {
	//	st = jen.Op("*").Add(st)
	//}

	switch v := expr.(type) {
	case *ast.Ident:
		st.Id(v.Name)
	case *ast.SelectorExpr:
		x, ok := v.X.(*ast.Ident)
		if !ok {
			fmt.Printf("Unknown type: %#v.\n", x)
			return nil, 0
		}
		imp := x.Name
		if v, ok := imports[imp]; ok {
			imp = v
		}
		if v.Sel == nil {
			fmt.Printf("Empty name for type %v.\n", expr)
			return nil, 0
		}
		st.Qual(imp, v.Sel.Name)
	case *ast.StructType:
		st.Id("struct").Op("{").Line()
		for _, field := range v.Fields.List {
			fst, lvl := buildTypeDeclaration(field.Type, imports)
			if len(field.Names) > 1 {
				fmt.Println("Filed has more than one name. Please report this case.")
				return nil, 0
			}
			if len(field.Names) == 1 {
				st.Id(field.Names[0].Name)
			}
			st.Add(applyPointerLevel(fst, lvl))
			if field.Tag != nil {
				st.Id(field.Tag.Value)
			}
			st.Line()
		}
		st.Op("}")
	case *ast.MapType:
		key, lvl := buildTypeDeclaration(v.Key, imports)
		if key == nil {
			return nil, 0
		}
		value, lvl := buildTypeDeclaration(v.Value, imports)
		if value == nil {
			return nil, 0
		}
		st.Id("map").Op("[").Add(applyPointerLevel(key, lvl)).Op("]").Add(applyPointerLevel(value, lvl))
	case *ast.ArrayType:
		tp, lvl := buildTypeDeclaration(v.Elt, imports)
		st.Op("[").Op("]").Add(applyPointerLevel(tp, lvl))
	default:
		fmt.Printf("Unknown type: %v.\n", expr)
		return nil, 0
	}

	return st, lvl
}

func applyPointerLevel(stmt *jen.Statement, level int) *jen.Statement {
	op := jen.Empty()
	for i := 0; i < level; i++ {
		op.Op("*")
	}
	return op.Add(stmt)
}
