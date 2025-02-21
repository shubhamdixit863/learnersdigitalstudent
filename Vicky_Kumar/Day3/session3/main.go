package main

import (
	"fmt"
)

func main() {
	// // typr inference
	// c := 9
	// fmt.Println(reflect.TypeOf(c))

	// // type conversion
	// // converting from one type to another
	// var d int
	// d = 9
	// f := float32(d)
	// fmt.Println(reflect.TypeOf(f))

	// var g float32
	// g = 9.78
	// i := int(g)

	// fmt.Println(i)

	// in, err := strconv.Atoi("34")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(in)

	// j := strconv.Itoa(2)
	// fmt.Println(reflect.TypeOf(j))

	// char := 'A'
	// jl := int(char)
	// fmt.Println("ascii to int:", jl)

	// Pass by value and pass by refrence:
	//go is pass by value

	slc := []int{2, 3}
	SliceModifier(slc)
	//[slc is created in stack] but underlying array is created in heap memory
	// value of slc is changed! why? bcoz sclice stores the address of the underlying array and when it is passed to the function the addres is copied to the function which still points to the same array.
	fmt.Println(slc)

	slc2 := [2]int{2, 3}
	SliceModifier2(slc2)
	// [array is created in stack] why? coz it has limited size
	// value of slc is not changed! why? bcoz array is copied so any change will not be reflect to array in main function
	fmt.Println(slc)

}

func SliceModifier(a []int) {
	a[0] = 9
}
func SliceModifier2(a [2]int) {
	a[0] = 9
}
