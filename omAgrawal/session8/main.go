package main

import (
	"encoding/json"
	"fmt"
	"session8/filer"
)

type Student struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {

	d := &Student{}

	data := []byte(`{"name":"Om","age":"20"}`)

	err := json.Unmarshal(data, d)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
	filer.ReadFile("./text.txt")

}
