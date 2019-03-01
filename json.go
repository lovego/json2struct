package json

import (
	"errors"
	"reflect"
)

var caches = make(map[reflect.Type]reflect.Type)

func UnmarshalToStruct(data []byte, p interface{}) ([]string, error) {
	return nil, nil
}

func getPtrFieldsStruct(p interface{}) (reflect.Type, error) {
	ptr := reflect.ValueOf(p)
	typ := ptr.Type()
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct || ptr.IsNil() {
		return nil, errors.New("arg p should be a non nil pointer to a struct")
	}
	return convertToPtrFieldsStruct(typ.Elem()), nil
}

func traverseStructFields(typ reflect.Type, fn func(field reflect.StructField)) bool {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		// exported field has an empty PkgPath
		if (!field.Anonymous || !traverseStructFields(field.Type, fn)) && field.PkgPath == "" {
			fn(field)
		}
	}
	return true
}
