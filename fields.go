package json2struct

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/lovego/structs"
)

var caches = make(map[reflect.Type][]fieldT)
var mutex = sync.RWMutex{}

type fieldT struct {
	Name, JsonKey string
}

func GetFields(p interface{}) ([]fieldT, error) {
	typ := reflect.TypeOf(p)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("arg p should be a struct or struct pointer")
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

	structs.TraverseType(typ, func(field reflect.StructField) bool {
		return field.Tag.Get("json") == "-"
	}, func(field reflect.StructField) {
		if key := getJSONKey(field.Name, field.Tag.Get("json")); key != "" {
			lower := strings.ToLower(key)
			fields = append(fields, fieldT{Name: field.Name, JsonKey: lower})
			m1[lower] = append(m1[lower], fieldT{Name: field.Name, JsonKey: key})
			m2[field.Name] = append(m2[field.Name], fieldT{Name: field.Name, JsonKey: key})
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

func getJSONKey(fieldName, tag string) string {
	if idx := strings.Index(tag, ","); idx != -1 {
		tag = tag[:idx]
	}
	if tag == "" {
		return fieldName
	}
	return tag
}
