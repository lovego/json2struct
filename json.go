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
	return typ.Elem(), nil
}
