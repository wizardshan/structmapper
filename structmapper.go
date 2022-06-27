package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"unicode"
)

var toName *string
var toPath *string
var fromName *string
var fromPath *string
var toVar string
var fromVar string
var funcName string
var funcPosition string

func init() {

	toName = flag.String("toName", "", "toName required")
	toPath = flag.String("toPath", "", "toPath required")
	fromName = flag.String("fromName", "", "fromName required")
	fromPath = flag.String("fromPath", "", "fromPath required")

	flag.StringVar(&toVar, "toVar", "", "toVar not required")
	flag.StringVar(&fromVar, "fromVar", "", "fromVar not required")

	flag.StringVar(&funcPosition, "funcPosition", "", "default toObject")
}

type Object struct {
	Info    *ast.Object
	Name    string
	PkgName string
	Path    string

	FuncMapping map[string]*ast.FuncDecl
}

type ObjectMapping map[string]*Object

func (mapping ObjectMapping) Fetch(name string) *Object {
	if mapping.Exit(name) {
		return mapping[name]
	}

	o := new(Object)
	o.Name = name
	o.FuncMapping = make(map[string]*ast.FuncDecl)
	mapping[name] = o
	return o
}

func (mapping ObjectMapping) Exit(name string) bool {
	if _, ok := mapping[name]; ok {
		return true
	}

	return false
}

type ObjectInfo struct {
	Name string
	Path string
}

const (
	funcPositionTo = "to"
	funcPositionFrom = "from"
)

var toObjectMapping = make(ObjectMapping)
var fromObjectMapping = make(ObjectMapping)

func main() {
	flag.Parse()

	toInfo := new(ObjectInfo)
	toInfo.Name = *toName
	toInfo.Path = *toPath

	fromInfo := new(ObjectInfo)
	fromInfo.Name = *fromName
	fromInfo.Path = *fromPath

	parseFile(toObjectMapping, toInfo.Path)
	parseFile(fromObjectMapping, fromInfo.Path)

	if !toObjectMapping.Exit(toInfo.Name) {
		log.Fatalf("toObject not exit")
	}

	toObject := toObjectMapping[toInfo.Name]
	toObject.Path = toInfo.Path

	if !fromObjectMapping.Exit(fromInfo.Name) {
		log.Fatalf("fromObject not exit")
	}

	fromObject := fromObjectMapping[fromInfo.Name]
	fromObject.Path = fromInfo.Path

	src := genString(toObject, fromObject, toVar, fromVar, funcPosition)
	fileNamePrefix := toInfo.Name
	filePath := toInfo.Path
	if funcPosition == funcPositionFrom {
		  fileNamePrefix = fromInfo.Name
		  filePath = fromInfo.Path
	}

	fileName := fmt.Sprintf("%s_mapping.go", fileNamePrefix)
	outputName := filepath.Join(".", strings.ToLower(fileName))

	err := ioutil.WriteFile(filePath + "/" + outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

func generate(tmp *string, toObject *Object, fromObject *Object, toVarName string, fromVarName string, funcPosition string) {
	if toObject == nil {
		return
	}

	toFieldList := getFieldList(toObject.Info.Decl)
	if toFieldList == nil {
		return
	}

	var fromFieldList *ast.FieldList
	if fromObject != nil {
		fromFieldList = getFieldList(fromObject.Info.Decl)
	}

	pkgName := ""
	if funcPosition == funcPositionFrom {
		 pkgName = toObject.PkgName + "."
	}

	for _, item := range toFieldList.List {
		toField, ok := parseField(item)
		if !ok {
			continue
		}

		toFieldName := getFieldName(toField.Names)

		var fromField *ast.Field
		if fromFieldList != nil {
			fromField = getFieldByName(fromFieldList.List, toFieldName)
		}

		toFieldType := getType(toField.Type)

		//// 匿名
		//if toFieldType.Anonymous {
		//	toFieldType.Name = toFieldName
		//}

		var fromFieldType *Type
		if fromField != nil {
			fromFieldType = getType(fromField.Type)
		}

		toFieldObject := toObjectMapping[toFieldType.Name]
		fromFieldObject := fromObjectMapping[toFieldType.Name]

		var toFieldObjectType *Type
		if toFieldObject != nil {
			toFieldObjectType = getType(toFieldObject.Info.Decl)
		}

		var fromFunc *ast.FuncDecl
		if fromObject != nil {
			fromFunc = fromObject.FuncMapping[toFieldName]
		}

		if !isUpper(toFieldName) {
			*tmp += fmt.Sprintf("//%s.%s = private property \n", toVarName, toFieldName)
			continue
		}

		if toFieldObjectType != nil && !toFieldObjectType.IsObject {
			if fromField == nil && fromFunc == nil {
				goto FromFieldNotExist
			}

			parentheses := ""
			if fromFunc != nil {
				parentheses = "()"
			}

			fromStar := ""
			if fromFieldType.Star {
				fromStar = "*"
				*tmp += fmt.Sprintf("if %s.%s != nil { \n", fromVarName, toFieldName)
			}

			if toFieldType.Star {
				varName := LowerFirst(toFieldName)
				*tmp += fmt.Sprintf("%s := %s%s(%s%s.%s%s) \n", varName, pkgName, toFieldType.Name, fromStar, fromVarName, toFieldName, parentheses)
				*tmp += fmt.Sprintf("%s.%s = &(%s) \n", toVarName, toFieldName, varName)
			} else {
				*tmp += fmt.Sprintf("%s.%s = %s%s(%s%s.%s%s) \n", toVarName, toFieldName, pkgName, toFieldType.Name, fromStar, fromVarName, toFieldName, parentheses)
			}

			if fromFieldType.Star {
				*tmp += fmt.Sprintf("}\n\n")
			}

			continue
		}

		if fromFunc != nil {
			*tmp += fmt.Sprintf("%s.%s = %s.%s() \n", toVarName, toFieldName, fromVarName, toFieldName)
			continue
		}

	FromFieldNotExist:
		if fromField == nil {
			*tmp += fmt.Sprintf("//%s.%s = fromStruct property not exist \n", toVarName, toFieldName)
			continue
		}

		symbol := ""
		if toFieldType.Star != fromFieldType.Star {
			symbol = "*"
			if toFieldType.Star {
				symbol = "&"
			}
		}

		if !toFieldType.IsObject {
			if fromFieldType.Star {
				*tmp += fmt.Sprintf("if %s.%s != nil { \n", fromVarName, toFieldName)
			}
			*tmp += fmt.Sprintf("%s.%s = %s%s.%s \n", toVarName, toFieldName, symbol, fromVarName, toFieldName)
			if fromFieldType.Star {
				*tmp += fmt.Sprintf("}\n\n")
			}
			continue
		}

		*tmp += "\n"

		if !toFieldType.IsSlice && (toFieldObjectType != nil && !toFieldObjectType.IsSlice) {

			if fromFieldType.Star {
				*tmp += fmt.Sprintf("if %s.%s != nil { \n", fromVarName, toFieldName)
				*tmp += fmt.Sprintf("%s := new(%s%s) \n", toVarName+toFieldName, pkgName, toFieldType.Name)
			} else {
				*tmp += fmt.Sprintf("var %s %s%s \n", toVarName+toFieldName, pkgName, toFieldType.Name)
			}

			generate(tmp, toFieldObject, fromFieldObject, toVarName+toFieldName, fromVarName+"."+toFieldName, funcPosition)

			*tmp += fmt.Sprintf("%s.%s = %s%s \n", toVarName, toFieldName, symbol, toVarName+toFieldName)

			if fromFieldType.Star {
				*tmp += fmt.Sprintf("}\n\n")
			}

			continue
		}

		varName := strings.ReplaceAll(fromVarName, ".", "")
		lenName := varName + toFieldName + "Len"
		itemName := varName + toFieldName + "Item"
		indexName := varName + toFieldName + "Index"

		toStar := ""
		if toFieldType.Star {
			toStar = "*"
		}

		varSliceName := toVarName + toFieldName

		brackets := "[]"
		itemTypeName := toFieldType.Name
		varItemName := toVarName + toFieldType.Name
		if toFieldObjectType != nil && toFieldObjectType.IsSlice {
			brackets = ""
			itemTypeName = toFieldObjectType.Name
			varItemName = toVarName + toFieldObjectType.Name
		}

		*tmp += fmt.Sprintf("%s := len(%s.%s) \n", lenName, fromVarName, toFieldName)
		*tmp += fmt.Sprintf("%s := make(%s%s%s%s, %s) \n", varSliceName, brackets, toStar, pkgName, toFieldType.Name, lenName)
		*tmp += fmt.Sprintf("if %s > 0 { \n", lenName)
		*tmp += fmt.Sprintf("for %s := 0; %s < %s; %s++  {\n", indexName, indexName, lenName, indexName)
		*tmp += fmt.Sprintf("%s := %s.%s[%s] \n", itemName, fromVarName, toFieldName, indexName)
		*tmp += fmt.Sprintf("%s := new(%s%s) \n", varItemName, pkgName, itemTypeName)

		if !toObjectMapping.Exit(toFieldType.Name) {
			log.Fatalf("toObject toFieldTypeName %s not exit", toFieldType.Name)
		}

		if toFieldObjectType != nil && toFieldObjectType.IsSlice {
			toFieldObject = toObjectMapping[toFieldObjectType.Name]
			fromFieldObject = fromObjectMapping[toFieldObjectType.Name]
		}

		generate(tmp, toFieldObject, fromFieldObject, varItemName, itemName, funcPosition)

		*tmp += fmt.Sprintf("%s[%s] = %s \n", varSliceName, indexName, varItemName)
		*tmp += "}\n"
		*tmp += "}\n\n"
		*tmp += fmt.Sprintf("%s.%s = %s \n", toVarName, toFieldName, toVarName+toFieldName)
	}

}

func genString(toObject *Object, fromObject *Object, toVarName string, fromVarName string, funcPosition string) []byte {
	if toVarName == "" {
		toVarName = toObject.PkgName
	}

	if fromVarName == "" {
		fromVarName = fromObject.PkgName
	}

	toObjectType := getType(toObject.Info.Decl)
	var toObjectIsSlice bool
	symbol := "*"
	if toObjectType.IsObject && toObjectType.IsSlice {
		toObjectIsSlice = true
		symbol = ""
	}

	packageName := toObject.PkgName
	pkgName := ""
	if funcPosition == funcPositionFrom {
		packageName = fromObject.PkgName
		pkgName = toObject.PkgName + "."
	}

	tmp := fmt.Sprintf("package %s", packageName)
	tmp += "\n"
	//tmp += fmt.Sprintf(`import "%s/%s"`, moduleName, fromObject.Path)
	tmp += "\n"

	if funcPosition == funcPositionFrom {
		tmp += fmt.Sprintf(`func (%s %s%s) Mapping() %s%s.%s  {`, fromVarName, symbol, fromObject.Name, symbol, toObject.PkgName, toObject.Name)
	} else {
		tmp += fmt.Sprintf(`func (%s *%s) Mapping(%s %s%s.%s) {`, toVarName, toObject.Name, fromVarName, symbol, fromObject.PkgName, fromObject.Name)
	}

	tmp += "\n"
	tmp += "/**************** mapping start ****************/ \n"

	if toObjectIsSlice {

		toFieldName := toObject.Name
		varName := strings.ReplaceAll(fromVarName, ".", "")
		lenName := varName + toFieldName + "Len"
		itemName := varName + toFieldName + "Item"
		indexName := varName + toFieldName + "Index"

		varItemName := toVarName + toObjectType.Name
		itemTypeName := toObjectType.Name

		tmp += fmt.Sprintf("%s := len(%s) \n", lenName, fromVarName)
		if funcPosition == funcPositionFrom {
			tmp += fmt.Sprintf("%s := make(%s%s, %s) \n", toVarName, pkgName, toFieldName, lenName)
		} else {
			tmp += fmt.Sprintf("*%s = make(%s, %s) \n", toVarName, toFieldName, lenName)
		}
		tmp += fmt.Sprintf("if %s > 0 { \n", lenName)

		tmp += fmt.Sprintf("for %s := 0; %s < %s; %s++  {\n", indexName, indexName, lenName, indexName)
		tmp += fmt.Sprintf("%s := %s[%s] \n", itemName, fromVarName, indexName)
		tmp += fmt.Sprintf("%s := new(%s%s) \n", varItemName, pkgName, itemTypeName)

		toObject = toObjectMapping[toObjectType.Name]
		fromObject = fromObjectMapping[toObjectType.Name]

		generate(&tmp, toObject, fromObject, varItemName, itemName, funcPosition)

		if funcPosition == funcPositionFrom {
			tmp += fmt.Sprintf("%s[%s] = %s \n", toVarName, indexName, varItemName)
		} else {
			tmp += fmt.Sprintf("(*%s)[%s] = %s \n", toVarName, indexName, varItemName)
		}

		tmp += "}\n"

		if funcPosition == funcPositionFrom {
			tmp += fmt.Sprintf("return %s \n", toVarName)
		}

		tmp += "}\n\n"
	} else {
		if funcPosition == funcPositionFrom {
			tmp += fmt.Sprintf("%s := new (%s.%s) \n", toVarName, toObject.PkgName, toObject.Name)
		}

		generate(&tmp, toObject, fromObject, toVarName, fromVarName, funcPosition)
	}
	if funcPosition == funcPositionFrom {
		if toObjectIsSlice {
			tmp += fmt.Sprintf("return nil")
		} else {
			tmp += fmt.Sprintf("return %s \n", toVarName)
		}
	}

	tmp += "\n/**************** mapping end  ****************/"
	tmp += "}"

	buff := bytes.NewBufferString("")
	buff.WriteString(tmp)

	//格式化
	src, err := format.Source(buff.Bytes())
	if err != nil {
		log.Fatalf("format.Source %v", err)
	}
	return src
}

type Type struct {
	Name     string
	Star     bool
	IsObject bool
	IsArray  bool
	IsSlice  bool
	Len      int
	//Anonymous bool
}

func getType(i interface{}) *Type {

	if i == nil {
		return nil
	}

	t := new(Type)

	typeSpec, ok := parseTypeSpec(i)
	if ok {
		i = typeSpec.Type
	}

	arrayType, ok := parseArrayType(i)
	if ok {
		if arrayType.Len != nil {
			t.IsArray = true
		} else {
			t.IsSlice = true
		}
		i = arrayType.Elt
	}

	starExpr, ok := parseStarExpr(i)
	if ok {
		t.Star = true
		i = starExpr.X
	}

	selectorExpr, ok := parseSelectorExpr(i)
	if ok {
		i = selectorExpr.X
	}

	ident, ok := parseIdent(i)
	if ok {
		t.Name = ident.Name
	}

	//_, ok = parseStructType(i)
	//if ok {
	//	t.Anonymous = true
	//}

	if isObjectType(t.Name) {
		t.IsObject = true
	}

	return t
}

type FuncInfo struct {
	Name    string
	SelStar bool
	SelName string
}

func getFuncInfo(funcDecl *ast.FuncDecl) *FuncInfo {

	if funcDecl.Recv == nil {
		return nil
	}

	info := new(FuncInfo)
	ident, ok := parseIdent(funcDecl.Name)

	if ok {
		info.Name = ident.Name
	}

	for _, field := range funcDecl.Recv.List {
		t := getType(field.Type)
		info.SelName = t.Name
		info.SelStar = t.Star
	}

	return info
}

func parseTypeSpec(i interface{}) (*ast.TypeSpec, bool) {
	n, ok := i.(*ast.TypeSpec)
	return n, ok
}

func parseStructType(i interface{}) (*ast.StructType, bool) {
	n, ok := i.(*ast.StructType)
	return n, ok
}

func parseFieldList(i interface{}) (*ast.FieldList, bool) {
	n, ok := i.(*ast.FieldList)
	return n, ok
}

func parseField(i interface{}) (*ast.Field, bool) {
	n, ok := i.(*ast.Field)
	return n, ok
}

func parseStarExpr(i interface{}) (*ast.StarExpr, bool) {
	n, ok := i.(*ast.StarExpr)
	return n, ok
}

func parseSelectorExpr(i interface{}) (*ast.SelectorExpr, bool) {
	n, ok := i.(*ast.SelectorExpr)
	return n, ok
}

func parseArrayType(i interface{}) (*ast.ArrayType, bool) {
	n, ok := i.(*ast.ArrayType)
	return n, ok
}

func parseIdent(i interface{}) (*ast.Ident, bool) {
	n, ok := i.(*ast.Ident)
	return n, ok
}

func getFieldList(i interface{}) *ast.FieldList {
	typeSpec, ok := parseTypeSpec(i)
	if !ok {
		return nil
	}

	structType, ok := parseStructType(typeSpec.Type)
	if !ok {
		return nil
	}

	fieldList, ok := parseFieldList(structType.Fields)
	if !ok {
		return nil
	}

	return fieldList
}

func getFieldName(names []*ast.Ident) string {
	for _, item := range names {
		return item.Name
	}
	return ""
}

func getFieldByName(fields []*ast.Field, name string) *ast.Field {
	for _, item := range fields {
		field, ok := parseField(item)
		if !ok {
			continue
		}

		if name == getFieldName(field.Names) {
			return field
		}
	}

	return nil
}

func parseFile(objectMapping ObjectMapping, path string) {

	path = "./" + path + "/"
	// 解析当前目录下包信息
	pkgInfo, err := build.ImportDir(path, 0)
	if err != nil {
		log.Fatalf("path = %s build.ImportDir %v", path, err)
	}

	fset := token.NewFileSet()
	for _, file := range pkgInfo.GoFiles {
		f, err := parser.ParseFile(fset, path+file, nil, 0)
		if err != nil {
			log.Fatalf("parser.ParseFile %v", err)
		}

		ast.Inspect(f, func(n ast.Node) bool {
			file, ok := n.(*ast.File)
			if !ok {
				return false
			}

			pkgName := file.Name.String()

			for _, decl := range f.Decls {
				funcDecl, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}

				funcInfo := getFuncInfo(funcDecl)
				if funcInfo == nil {
					continue
				}

				o := objectMapping.Fetch(funcInfo.SelName)
				o.PkgName = pkgName
				o.FuncMapping[funcInfo.Name] = funcDecl

			}

			for k, obj := range file.Scope.Objects {

				o := objectMapping.Fetch(k)
				o.PkgName = pkgName
				o.Info = obj

				// 匿名struct
				fieldList := getFieldList(obj.Decl)
				if fieldList == nil {
					continue
				}

				for _, field := range fieldList.List {
					structType, ok := parseStructType(field.Type)
					if !ok {
						continue
					}

					fieldO := new(ast.Object)
					fieldO.Name = getFieldName(field.Names)
					fieldO.Decl = &ast.TypeSpec{
						Type:structType,
					}

					fieldObject := objectMapping.Fetch(fieldO.Name)
					fieldObject.PkgName = pkgName
					fieldObject.Info = fieldO

				}
			}

			return true
		})
	}
}

func isObjectType(dataType string) bool {
	dataTypes := map[string]bool{
		"uint8":      true,
		"uint16":     true,
		"uint32":     true,
		"uint64":     true,
		"int8":       true,
		"int16":      true,
		"int32":      true,
		"int64":      true,
		"int":        true,
		"float32":    true,
		"float64":    true,
		"complex64":  true,
		"complex128": true,
		"byte":       true,
		"rune":       true,
		"uintptr":    true,
		"string":     true,
		"time":       true,
		"slice":      true,
	}

	if _, ok := dataTypes[dataType]; ok {
		return false
	}
	return true
}

func isUpper(str string) bool {
	return unicode.IsUpper(rune(str[0]))
}

func LowerFirst(str string) string {
	return string(unicode.ToLower(rune(str[0]))) + str[1:]

}