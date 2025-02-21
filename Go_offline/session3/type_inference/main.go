package main

import (
	"fmt"
	"reflect" // for type of variable
	"strconv" // for string to int conversion
)

func main() {
	//type inference
	c:=9 //now c is automatically infered as int
	fmt.Println(reflect.TypeOf(c))


	fmt.Println("-----------------")



	//type conversion int to float

	var a int = 10
	var b float64 = float64(a) //explicit conversion or type casting
	fmt.Println(reflect.TypeOf(b))
	fmt.Println("-----------------")


	//type conversion flaoat to int
	var d float32 = 10.67
	var e int = int(d) //explicit conversion
	fmt.Println(reflect.TypeOf(e))
	fmt.Println("-----------------")

	//type conversion bool to int
	var f bool = true
	var g int
	if f {
		g = 1
	} else {
		g = 0
	}
	fmt.Println(reflect.TypeOf(g))
	fmt.Println("-----------------")
	
   //string to int
	var h string = "10"
	i, err:= strconv.Atoi(h)
	if err != nil {
		fmt.Println("Conversion not possible", err)
	}

	fmt.Println(i)

	var s string = "A"
	Integer, err:= strconv.Atoi(s)
	if err != nil {
		fmt.Println("Conversion not possible", err)
	}

	fmt.Println(Integer)
	fmt.Println("-----------------")


	//int to string
	var num int = 10
	var str string = strconv.Itoa(num)
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))
	fmt.Println("-----------------")


	//ascii to int
	char:= 'A'
	fmt.Println(reflect.TypeOf(char))
	Int_value:= int(char)
	fmt.Println("Value of data type is", Int_value)

}