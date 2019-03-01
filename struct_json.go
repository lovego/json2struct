package json

import (
	"errors"
	"reflect"
	"strings"
)

type structJson struct {
	keyToField map[string]string
}

func newStructJson(typ reflect.Type) (*structJson, error) {
	m := make(map[string]string)
	var err error
	traverseStructFields(typ, func(field reflect.StructField) {
		if key := getJsonKey(field.Name, field.Tag.Get("json")); key != "" {
			lower := strings.ToLower(key)
			if _, ok := m[lower]; ok {
				err = errors.New("conflict json key: " + key)
			} else {
				m[lower] = field.Name
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return &structJson{keyToField: m}, nil
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
		// exported field has an empty PkgPath
		if (!field.Anonymous || !traverseStructFields(field.Type, fn)) && field.PkgPath == "" {
			fn(field)
		}
	}
	return true
}
