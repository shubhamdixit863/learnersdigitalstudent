package main

import (
	"fmt"
	"strings"
)

func Add(i interface{}) interface{} {
	// num, ok := i.(int)
	// if ok {
	// 	return num * 2
	// }
	// //type assertion
	// num2, ok := i.(float64)
	// if ok {
	// 	return num2 * 2
	// }
	// return nil

	switch v := i.(type) {
	case float64:
		return v
	case string:
		return strings.ToLower(v)
	default:
		return fmt.Sprintf("%v", v)
	}

}
func main() {
	b := true
	fmt.Println(Add(9))
	fmt.Println(Add(9.9))
	fmt.Println(Add("345"))
	fmt.Println(Add(b))
}
