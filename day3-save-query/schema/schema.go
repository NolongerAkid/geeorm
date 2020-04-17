package schema

import (
	"awesomeProject4/geeorm/day2-reflect-schema/dialect"
	"go/ast"
	"reflect"
)

//Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag string
}

//Schema represents a table of database
type Schema struct {
	Model interface{}//被映射的对象
	Name string//表名
	Fields []*Field
	FieldNames []string//所有字段名（列名）
	fieldMap map[string]*Field
}

func (schema *Schema) GetField(name string) *Field{
	return schema.fieldMap[name]
}

func(schema *Schema) RecordValues(dest interface{})[]interface{}{
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _,field := range schema.Fields{
		fieldValues = append(fieldValues,destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

func Parse(dest interface{},d dialect.Dialect) *Schema{
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:      dest,
		Name:       modelType.Name(),
		Fields:     nil,
		FieldNames: nil,
		fieldMap:   make(map[string]*Field),
	}

	for i := 0;i < modelType.NumField(); i++{
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name){
			field := &Field{
				Name : p.Name,
				Type : d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v,ok := p.Tag.Lookup("geeorm");ok{
				field.Tag = v
			}
			schema.Fields = append(schema.Fields,field)
			schema.FieldNames = append(schema.FieldNames,p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}