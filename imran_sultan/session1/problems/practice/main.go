package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello")
	// a, b := swap(5, 2)
	// println(a, b)
	// area(4)
	// celciusToferenhite(32)
	// concatStr("imran", "ali")
	// checkVariable(1, 2)
	varadicFunc(1, 2, 3, 4)

}
func swap(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b

}
func area(r float64) {
	var area = math.Pi * r * r
	fmt.Println(area)

}
func celciusToferenhite(c float32) {
	var f = (c * (9 / 5)) + 32
	fmt.Println(f)
}
func concatStr(s1 string, s2 string) {
	println(s1 + s2)

}
func checkVariable(a int, b int) {
	fmt.Printf("type of a is %T %T", a, b)

}
func varadicFunc(n ...int) {
	var sum = 0
	for i := 0; i < len(n); i++ {
		sum = sum + n[i]
	}
	fmt.Println(sum)
}
