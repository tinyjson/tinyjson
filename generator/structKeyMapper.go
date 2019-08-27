package generator

import (
	"fmt"
	"go/ast"
	"reflect"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
)

type (
	structKeyMapper struct {
		queue []structKeyMapperQueue
		dic   map[string]structKeyMapperKey
	}
	structKeyMapperKey struct {
		variableName *jen.Statement
		nestedLevel  int
		ignore       bool
		omitempty    bool
		typeExpr     ast.Expr
	}
	structKeyMapperQueue struct {
		fl       *ast.FieldList
		variable *jen.Statement
	}
)

func newStructKeyMapper(fl *ast.FieldList, variable *jen.Statement) *structKeyMapper {
	return &structKeyMapper{
		queue: []structKeyMapperQueue{{fl: fl, variable: variable}},
		dic:   map[string]structKeyMapperKey{},
	}
}

func (km *structKeyMapper) Proc(meta *TypeMeta) {
	var toNextLvl = 1
	var lvl int
	for {
		if len(km.queue) == 0 {
			break
		}
		if toNextLvl == 0 {
			toNextLvl = len(km.queue)
			lvl++
		}
		toNextLvl--
		el := km.queue[0]
		km.queue = km.queue[1:]

		for _, f := range el.fl.List {
			if len(f.Names) > 1 {
				fmt.Println("interesting case! please report it")
				continue
			}

			var (
				name      string
				ignore    bool
				omitempty bool
			)

			if f.Tag != nil {
				tag := reflect.StructTag(strings.TrimPrefix(strings.TrimSuffix(f.Tag.Value, "`"), "`"))
				name, ignore, omitempty = parseTag(tag)
				if ignore {
					continue
				}
			}

			var key structKeyMapperKey
			key.omitempty = omitempty
			key.typeExpr = f.Type
			key.nestedLevel = lvl

			if len(f.Names) == 0 {
				var pkg, typeName string
				switch v := f.Type.(type) {
				case *ast.Ident:
					if !v.IsExported() {
						continue
					}
					typeName = v.Name
				case *ast.SelectorExpr:
					x := v.X.(*ast.Ident)
					if x == nil {
						continue
					}
					pkg = meta.imports[x.Name]
					typeName = v.Sel.Name
				}
				key.variableName = el.variable.Clone().Dot(typeName)

				typeMeta := MetaRel(meta.filePath, pkg, typeName)
				if typeMeta == nil {
					fmt.Println("meta not found", meta.filePath, pkg, typeName)
				}

				if name == "" {
					if st, ok := typeMeta.spec.Type.(*ast.StructType); ok {
						km.queue = append(km.queue, structKeyMapperQueue{
							fl:       st.Fields,
							variable: el.variable.Clone().Dot(typeName),
						})
						continue
					} else if typeName == "" {
						continue
					} else {
						name = typeName
					}
				}
			} else {
				if name == "" {
					name = f.Names[0].Name
				}
				key.variableName = el.variable.Clone().Dot(f.Names[0].Name)
			}
			km.AddField(name, key)
		}
	}
	km.RemoveIgnored()
}

func (km *structKeyMapper) AddField(fieldName string, key structKeyMapperKey) {
	res, ok := km.dic[fieldName]
	if ok {
		if res.nestedLevel == key.nestedLevel {
			res.ignore = true
			km.dic[fieldName] = res
		} // otherwise nested level is greater
	} else {
		km.dic[fieldName] = key
	}
}

func (km *structKeyMapper) RemoveIgnored() {
	for k, v := range km.dic {
		if v.ignore {
			delete(km.dic, k)
		}
	}
}

func (km *structKeyMapper) Keys() []string {
	var keys []string
	for k := range km.dic {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func parseTag(tag reflect.StructTag) (name string, ignore bool, omitempty bool) {
	val := strings.Split(tag.Get("json"), ",")
	if len(val) > 1 && val[1] == "omitempty" {
		omitempty = true
	}
	if val[0] == "-" {
		ignore = true
	}
	if val[0] != "" {
		name = val[0]
	}
	return
}
