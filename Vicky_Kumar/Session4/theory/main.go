package main

import "fmt"

func main() {
	// map is a descriptor
	// p := make(map[string]int) // non-nil
	// p["name"] = 2
	// p["john"] = 3
	// fmt.Println(p)

	// student class room example student ->no seat-> panic
	var m map[string]int //nil map -> no storage : underlying hash table is not created if used this syntax
	fmt.Println(m == nil)
	fmt.Println(m["hello"])
	m["hello"] = 2

	var mp = map[string]int{
		"name": 1,
		"john": 2,
	}
	fmt.Println(mp)

}
func Comparison() {
	// m1 := make(map[string]int)
	// m1["hello"] = 1
	// m2 := make(map[string]int)
	// fmt.Println(m1 == m2) //map can not compared with each other but can only be compared with nil
	// // same goes for slice but arrays can be compared
	// s1 := make([]int, 2)
	// s2 := make([]int, 2)
	// fmt.Println(s1 == s2)

	a1 := [2]int{1, 2}
	a2 := [2]int{1, 2}
	fmt.Println(a1 == a2)
}
