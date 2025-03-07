package main

import "fmt"

func main() {
	// c := square(8)
	// fmt.Println(c)
	// d := sum(12, 4)
	// fmt.Println(d)

	r, err := divide(10, 3)
	// always check the error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

	s, _ := divide(10, 3) // if you want to ignore a variable use _
	fmt.Println(s)

	VariadicFunc(6.4, 2, 4, 6, 8, 10)
}

func square(num int) int {
	c := num ^ 2
	return c
}
func sum(num1 int, num2 int) int {
	c := num1 + num2
	return c
}
func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("you can't divide by zero %s", "divide by zero")
	}
	c := num1 / num2
	return c, nil
}
func VariadicFunc(f float32, n ...int) []int { // variadic params need to be at the end
	// fmt.Println(f, n)
	// for loop
	// for i := 0; i < len(n); i++ {
	// 	fmt.Println(n[i])
	// }

	// n is slice
	return n
}

func namedReturn(n int) (double int) { //  easy to understand
	double = 2 * n
	return // no need to return double here
}
