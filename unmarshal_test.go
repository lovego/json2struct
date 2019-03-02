package json2struct

import (
	"encoding/json"
	"fmt"
)

func ExampleUnmarshal() {
	var v struct {
		A, B, C int64
	}
	fields, err := Unmarshal([]byte(`{"a": 0, "c": 2}`), &v)

	fmt.Printf("%+v\n", v)
	fmt.Println(fields, err)

	// Output:
	// {A:0 B:0 C:2}
	// [A C] <nil>
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
		Id   int
		Name string `json:"userName,omitempty"`
		typ1
	}{})

	testUnmarshal(testContent, &struct {
		Id   int
		Name string `json:"userName,omitempty"`
		typ1
	}{})

	testUnmarshal(testContent, &struct {
		Id   int
		Name string `json:"name"`
		typ1
		*typ2
		Typ2
	}{})

	// Output:
	// &{}
	// []
	// &{Id:0 Name: typ1:{Bank: Ignored:0 unexported:0}}
	// []
	// &{Id:1 Name:name typ1:{Bank:bank Ignored:0 unexported:0}}
	// [Id Name Bank]
	// &{Id:1 Name:name typ1:{Bank:bank Ignored:0 unexported:0} typ2:<nil> Typ2:Typ2}
	// [Id Name Bank Typ2]
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
		Id   int
		Name string `json:"userName,omitempty"`
		typ1
	}{})

	testAffected(testContent, struct {
		Id   int
		Name string `json:"userName,omitempty"`
		typ1
	}{})

	testAffected(testContent, struct {
		Id   int
		Name string `json:"name"`
		typ1
		typ2
		Typ2
	}{})

	// Output:
	// []
	// []
	// [Id Name Bank]
	// [Id Name Bank Typ2]
}

func testAffected(s string, p interface{}) {
	if fields, err := Affected([]byte(s), p); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fields)
	}
}

func Example_Unmarshal_omitempty() {
	// omitempty should not work with Unmarshal
	var v = struct {
		A string `json:",omitempty"`
	}{A: "A"}

	fmt.Println(v, json.Unmarshal([]byte(`{"A":""}`), &v), v)
	// Output:
	// {A} <nil> {}
}

func Example_Unmarshal_embedded() {
	type T int64
	var v struct{ T }

	fmt.Println(v, json.Unmarshal([]byte(`{"t":9}`), &v), v)
	// Output:
	// {0} <nil> {9}
}
