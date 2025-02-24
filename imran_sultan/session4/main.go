package main

import "fmt"

func fill(m map[string]int) {
	m["imran"] = 102
}

func main() {
	fmt.Println("hello world")
	// var my_slice_1 = make([]int, 4, 7)
	// fmt.Println(my_slice_1)
	// p := make(map[string]int)
	// fmt.Println(p == nil)
	// // fmt.Println(p["hello"])
	// p["name"] = 1
	// p["imran"] = 7
	// fmt.Println(p)
	// var m map[string]int

	// fmt.Println(m == nil)
	// fmt.Println(m["hee"])
	// fill(p)
	// fmt.Println(p["imran"])
	// var slice_1 = make([]int, 4, 7)
	// var slice_2 = make([]int, 4, 7)
	// // fmt.Println(slice_1 == slice_2)
	// m1 := make(map[string]int)
	// m2 := make(map[string]int)
	// // fmt.Println(m1 == m2)
	aar1 := [2]int{1, 2}
	aar2 := [2]int{1, 2}
	fmt.Println(aar1 == aar2)

}
