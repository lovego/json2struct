package json2struct

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var caches = make(map[reflect.Type][]fieldT)
var mutex = sync.RWMutex{}

type fieldT struct {
	name, jsonKey string
}

func getFields(p interface{}) ([]fieldT, error) {
	typ := reflect.TypeOf(p)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("arg p is not a struct")
	}
	mutex.RLock()
	fields := caches[typ]
	mutex.RUnlock()
	if fields == nil {
		var err error
		if fields, err = newFields(typ); err != nil {
			return nil, err
		}
		mutex.Lock()
		caches[typ] = fields
		mutex.Unlock()
	}
	return fields, nil
}

func newFields(typ reflect.Type) ([]fieldT, error) {
	var fields = make([]fieldT, 0, typ.NumField())
	var m1 = make(map[string][]fieldT)
	var m2 = make(map[string][]fieldT)

	traverseStructFields(typ, func(field reflect.StructField) {
		if key := getJsonKey(field.Name, field.Tag.Get("json")); key != "" {
			lower := strings.ToLower(key)
			fields = append(fields, fieldT{name: field.Name, jsonKey: lower})
			m1[lower] = append(m1[lower], fieldT{name: field.Name, jsonKey: key})
			m2[field.Name] = append(m2[field.Name], fieldT{name: field.Name, jsonKey: key})
		}
	})

	for _, conflicts := range m1 {
		if len(conflicts) > 1 {
			return nil, fmt.Errorf("conflicts field jsonKey: %+v", conflicts)
		}
	}
	for _, conflicts := range m2 {
		if len(conflicts) > 1 {
			return nil, fmt.Errorf("conflicts field name: %+v", conflicts)
		}
	}

	return fields, nil
}

func getJsonKey(fieldName, tag string) string {
	if tag == "-" {
		return ""
	}
	if idx := strings.Index(tag, ","); idx != -1 {
		tag = tag[:idx]
	}
	if tag == "" {
		return fieldName
	}
	return tag
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
		if (!field.Anonymous || !traverseStructFields(field.Type, fn)) &&
			(field.Name[0] >= 'A' && field.Name[0] <= 'Z') {
			fn(field)
		}
	}
	return true
}
