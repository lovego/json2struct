package json

import "reflect"

func convertToPtrFieldsStruct(typ reflect.Type) reflect.Type {
	var fields = make([]reflect.StructField, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fields[i] = convertToPtrField(typ.Field(i))
	}
	return reflect.StructOf(fields)
}

func convertToPtrField(field reflect.StructField) reflect.StructField {
	if field.Anonymous {
		switch field.Type.Kind() {
		case reflect.Struct:
			field.Type = convertToPtrFieldsStruct(field.Type)
			return field
		case reflect.Ptr:
			if elem := field.Type.Elem(); elem.Kind() == reflect.Struct {
				field.Type = reflect.PtrTo(convertToPtrFieldsStruct(elem))
				return field
			}
		}
	}
	// exported field has an empty PkgPath
	if field.PkgPath == "" {
		switch field.Type.Kind() {
		case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface,
			reflect.Chan, reflect.Func:
		default:
			field.Type = reflect.PtrTo(field.Type)
		}
	}
	return field
}
