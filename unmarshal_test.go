package json2struct

import (
	"encoding/json"
	"fmt"
)

func ExampleUnmarshal_omitempty() {
	// omitempty should not work with Unmarshal
	var v = struct {
		A string `json:",omitempty"`
	}{A: "A"}

	fmt.Println(v, json.Unmarshal([]byte(`{"A":""}`), &v), v)
	// Output:
	// {A} <nil> {}
}

func ExampleUnmarshal_embedded() {
	type T int64
	var v struct{ T }

	fmt.Println(v, json.Unmarshal([]byte(`{"t":9}`), &v), v)
	// Output:
	// {0} <nil> {9}
}
