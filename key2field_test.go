package json2struct

import (
	"fmt"
	"reflect"

	"github.com/lovego/maps"
)

func ExampleGetKey2FieldMap() {
	type typ1 struct {
		A          int64  `json:"-"`
		Name       string `json:"userName"`
		unexported int64
	}
	type Typ2 string
	type typ2 string

	for _, v := range []interface{}{
		struct{}{},
		struct {
			Id    int
			Name  string `json:"userName,omitempty"`
			other float32
		}{},
		struct {
			Id   int
			Name string `json:"name"`
			typ1
		}{},
		struct {
			Id   int
			Name string `json:"name"`
			*typ1
			*typ2
			Typ2
		}{},
		struct {
			Id   int
			Name string `json:"username"`
			typ1
		}{},
	} {
		if m, err := getKey2FieldMap(reflect.TypeOf(v)); err != nil {
			fmt.Println(err)
		} else {
			maps.Println(m)
		}
	}
	// Output:
	// map[]
	// map[id:Id username:Name]
	// map[id:Id name:Name username:Name]
	// map[id:Id name:Name typ2:Typ2 username:Name]
	// conflict json key: userName
}
