package json2struct

import (
	"fmt"
	"reflect"
)

func Example_NewFields() {
	for _, v := range []interface{}{
		struct{}{},
		struct {
			ID   int
			Name string `json:"userName,omitempty"`
			typ1
			other float32
		}{},
		struct {
			ID   int
			Name string `json:"name"`
			*typ1
			*typ2
			Typ2
		}{},
		struct {
			ID   int
			Name string `json:"bank"`
			typ1
		}{},
		struct {
			ID   int
			Bank string `json:"name"`
			typ1
		}{},
	} {
		if fields, err := newFields(reflect.TypeOf(v)); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fields)
		}
	}
	// Output:
	// []
	// [{ID id} {Name username} {Bank bank}]
	// [{ID id} {Name name} {Bank bank} {Typ2 typ2}]
	// conflicts field jsonKey: [{Name:Name JsonKey:bank} {Name:Bank JsonKey:bank}]
	// conflicts field name: [{Name:Bank JsonKey:name} {Name:Bank JsonKey:bank}]
}

func Example_GetFields() {
	type typ1 struct {
		A          int64  `json:"-"`
		Bank       string `json:"bank"`
		unexported int64
	}
	fmt.Println(GetFields(typ1{}))
	fmt.Println(GetFields(&typ1{}))
	// Output:
	// [{Bank bank}] <nil>
	// [{Bank bank}] <nil>
}
