package json2struct

import (
	"encoding/json"
	"fmt"
)

func ExampleUnmarshal() {
	var v struct {
		A, B, C int64
	}
	fields, err := Unmarshal([]byte(`{"a": 1, "b": 0}`), &v)

	fmt.Printf("%+v\n", v)
	fmt.Println(fields, err)

	// Output:
	// {A:1 B:0 C:0}
	// [A B] <nil>
}

type typ1 struct {
	Bank       string `json:"bank"`
	Ignored    int64  `json:"-"`
	unexported int64
}
type typ2 string
type Typ2 string

const testContent = `{
  "id":1, "Username": "name", "name": "name", "bAnK": "bank", "typ2": "type", "TYP2": "Typ2"
}`

func Example_Unmarshal() {
	testUnmarshal(testContent, &struct{}{})

	testUnmarshal("{}", &struct {
		ID   int
		Name string `json:"userName,omitempty"`
		typ1
	}{})

	testUnmarshal(testContent, &struct {
		ID   int
		Name string `json:"userName,omitempty"`
		typ1
	}{})

	testUnmarshal(testContent, &struct {
		ID   int
		Name string `json:"name"`
		typ1
		Typ2
	}{})

	// Output:
	// &{}
	// []
	// &{ID:0 Name: typ1:{Bank: Ignored:0 unexported:0}}
	// []
	// &{ID:1 Name:name typ1:{Bank:bank Ignored:0 unexported:0}}
	// [ID Name Bank]
	// &{ID:1 Name:name typ1:{Bank:bank Ignored:0 unexported:0} Typ2:Typ2}
	// [ID Name Bank Typ2]
}

func testUnmarshal(s string, p interface{}) {
	if fields, err := Unmarshal([]byte(s), p); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n%v\n", p, fields)
	}
}

func Example_Affected() {
	testAffected(testContent, struct{}{})

	testAffected("{}", struct {
		ID   int
		Name string `json:"userName,omitempty"`
		typ1
	}{})

	testAffected(testContent, struct {
		ID   int
		Name string `json:"userName,omitempty"`
		typ1
	}{})

	testAffected(testContent, struct {
		ID   int
		Name string `json:"name"`
		typ1
		typ2
		Typ2
	}{})

	// Output:
	// []
	// []
	// [ID Name Bank]
	// [ID Name Bank Typ2]
}

func testAffected(s string, p interface{}) {
	if fields, err := Affected([]byte(s), p); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fields)
	}
}

func Example_Unmarshal_omitempty() {
	// go1.12 and earlier versions omitempty tag make Unmarshal omit empty value.
	// go1.14 and later   versions omitempty tag has no effect on Unmarshal.
	var v = struct {
		A string `json:",omitempty"`
	}{A: "A"}

	fmt.Println(v, json.Unmarshal([]byte(`{"A":""}`), &v), v)
	// Output:
	// {} <nil> {}
}

func Example_Unmarshal_embedded() {
	type T int64
	var v struct{ T }

	fmt.Println(json.Unmarshal([]byte(`{"t":9}`), &v), v)
	// Output:
	// <nil> {9}
}
