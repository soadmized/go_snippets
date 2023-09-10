package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := MyStruct{
		Foo: 42,
		Bar: 777,
		Inner: InnerStruct{
			Smth:    "THIS",
			Another: 666,
		},
	}

	PrettyPrint(s)
}

type MyStruct struct {
	Foo   int         `json:"foo,omitempty"`
	Bar   int         `json:"bar,omitempty"`
	Inner InnerStruct `json:"inner"`
}

type InnerStruct struct {
	Smth    string `json:"smth,omitempty"`
	Another int    `json:"another,omitempty"`
}

func PrettyPrint(s interface{}) {
	structJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(structJSON))
}
