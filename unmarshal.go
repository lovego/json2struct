package json2struct

import (
	"encoding/json"
	"strings"
)

// Unmarshal decode data into a struct, and return the affected fields of the struct.
func Unmarshal(data []byte, p interface{}) ([]string, error) {
	if err := json.Unmarshal(data, p); err != nil {
		return nil, err
	}
	return Affected(data, p)
}

// Affected return the fields will be affected of a struct,
// if the data is decoded into the the struct.
func Affected(data []byte, p interface{}) ([]string, error) {
	fields, err := getFields(p)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	if len(m) == 0 {
		return nil, nil
	}

	m2 := make(map[string]struct{})
	for k := range m {
		m2[strings.ToLower(k)] = struct{}{}
	}
	var result []string
	for _, field := range fields {
		if _, ok := m2[field.jsonKey]; ok {
			result = append(result, field.name)
		}
	}
	return result, nil
}
