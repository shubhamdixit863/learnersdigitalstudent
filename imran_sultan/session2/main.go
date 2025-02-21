package main

import (
	"fmt"
	"services"
	"session2/services"
)

// func square(num int) int {
// 	c := num * num
// 	return c

// }

func main() {
	// fmt.Println("hi ia m a")
	// fmt.Println(square(2))
	// b := square(9)
	// fmt.Println(b)
	// fmt.Printf("%d", b)
	// fmt.Printf(b)
	// a1, b1 := divide(8, 0)
	// if b1 != nil {
	// 	fmt.Println(b1)
	// 	return
	// }
	// fmt.Println(a1, b1)
	var c = services.SayHello()
	println(c)
	// variedfunc(3.4, 12, 3, 3, 4, 45, 5)
	// fmt.Println(variedfunc(3.2, 1, 2, 3))
	// println(mysquart(2.3, false))
}

// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("you cannot divide by 0")
// 	}
// 	g := a / b
// 	return g, nil

// }
func variedfunc(f float32, n ...int) (float32, []int) {
	fmt.Println(f, n)
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
	return f, n
}
func mysquart(f1 float32, a bool) (f float32, ok bool) {
	f = f1 * 2
	ok = a
	return

}
