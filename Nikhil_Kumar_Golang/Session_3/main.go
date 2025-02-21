package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// type inference
		
	c := 9 // type inference now 

	// reflection
	fmt.Println("type of c: ", reflect.TypeOf(c))

	// type casting/conversion

	// int to float
	var a int = 9
	var b float64 = float64(a)
	fmt.Println(b)

	// float to int
	var x float64 = 9.9
	var y int = int(x)//it does not round off
	fmt.Println(y)

	// int to bool not possible
	// var p int = 1
	// var q bool = bool(p)
	// fmt.Println(q)

	// string to int
	var s string = "A"
	var t int = int(s[0])
	fmt.Println(t)

	// using strconv for type conversion
	// strinf to int
	h,err := strconv.Atoi("A")
	if err != nil {
		fmt.Println("conversion in not possible",err)
		//this will not give ascii value if number is in the string than it run otherwise not
	}else{
		fmt.Println(h)
	}
	// string to int

	j,err := strconv.Atoi("987987")
	if err != nil {
		fmt.Println("conversion in not possible",err)
	}else{
		fmt.Println(j)
	}

	// int to string
	var k int = 9
	var l string = strconv.Itoa(k)
	fmt.Println(l)

	// ascii to int
	char := 'A'
	m := int(char)
	fmt.Println(m) //


}