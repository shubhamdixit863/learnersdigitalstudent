package main

import (
	"fmt"
	"session9/internal/services"
	"strings"
)

func Add(i interface{}) interface{} {
	// Type Assertion
	// num := i.(int)

	// num, ok := i.(int)
	// if ok {
	// 	return num * 2
	// }

	// num2, ok := i.(float64)

	// if ok {
	// 	return num2 * 2
	// }

	// return nil

	switch v := i.(type) {
	case float32:
		return v

	case string:
		return strings.ToLower(v)

	default:
		return fmt.Sprintf("%v", v)
	}
}

func main() {
	var service services.Service
	service = services.NewEmployeeServiceV1()
	fmt.Println(service.GetData())

	Add(9)
	fmt.Println(Add(9))
	fmt.Println(Add(9.9))
}
