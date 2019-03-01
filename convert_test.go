package json

import (
	"fmt"
	"reflect"
)

func ExampleConvertToPtrFieldsStruct() {
	type t struct {
		D float64
	}
	for _, v := range []interface{}{
		struct{}{},
		struct{ A int }{},
		struct {
			A int
			B bool
		}{},
		struct {
			A int
			B bool
			C string
		}{},
		struct {
			A int
			B bool
			C string
			t
		}{},
	} {
		fmt.Println(convertToPtrFieldsStruct(reflect.TypeOf(v)))
	}
	// Output:
	// struct {}
	// struct { A *int }
	// struct { A *int; B *bool }
	// struct { A *int; B *bool; C *string }
}
