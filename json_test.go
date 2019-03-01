package json

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
