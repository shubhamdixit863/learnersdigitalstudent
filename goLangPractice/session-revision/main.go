package main

import (
	"fmt"
)

// const j = 9 // Can not change

// var (
// 	p = 10
// 	n = "hello"
// )

func main() {
	fmt.Println("Hello World")

	// var c string
	// fmt.Println(c)

	// h := 9
	// fmt.Println(h)

	// fmt.Println(p, n, j)

	// add.Sum()

	// var arr [6]int
	// fmt.Println(arr)

	// var slc []int
	// fmt.Println(slc)

	// slc2 := make([]int, 2) //make([]int, length, capacity(optional))
	// fmt.Println(slc2)
	// fmt.Println(len(slc2), cap(slc2))
	// fmt.Println(slc2)

	// slc2 = append(slc2, 9)
	// fmt.Println("Address of locker room: ", &slc2[0])

	// slc2 = append(slc2, 19)
	// fmt.Println("Address of locker room: ", &slc2[0])

	// slc2 = append(slc2, 190)
	// fmt.Println("Address of locker room: ", &slc2[0])

	// fmt.Println(slc2)
	// fmt.Println(len(slc2), cap(slc2))

	// mp := make(map[string]string)
	// mp["name"] = "sagar"
	// fmt.Println(mp)

	c := 9
	d := &c
	fmt.Println(d)
	fmt.Println(*d)

	foo(d)
	// foo(c)

	// a, b := foo2(d) // you can ignore value using underscore

	variadicfunction(9, 8, 7, 6)

	arr := []int{1, 2, 3, 4}

	s1 := arr[:2]

	s2 := arr[1:3]

	s1 = append(s1, 10) // Modify slice s1

	fmt.Println(arr, s1, s2)

}

func foo(d *int) int {
	return 10
}

func foo2(d *int) (int, string) {
	return 10, "ten"
}

func variadicfunction(n ...int) {
	fmt.Println(n)
}
