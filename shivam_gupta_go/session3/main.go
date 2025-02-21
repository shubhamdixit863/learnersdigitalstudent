package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
   c:=9
	 fmt.Println("value of c : ", c)
	 fmt.Println("type of c : ", reflect.TypeOf(c))
   // int to float 
	 b:= float32(c)
	 fmt.Println("value of b : ",b)
	 fmt.Println("type of b : ", reflect.TypeOf(b))
    
	 d:=9.34
	 fmt.Println("value of d : ", d)
	 fmt.Println("type of d : ",reflect.TypeOf(d))
  //float to int
	 a := int(d)
	 fmt.Println("value of a : ",a)
	 fmt.Println("type of a : ",reflect.TypeOf(a))

   //string to int
	 h,err:=strconv.Atoi("A")
	 if err!=nil {
		fmt.Println("conversion not possible ")
	 }else {fmt.Println(h)	 }
   
	 //string to int 
	 i,err := strconv.Atoi("123")
	 if err!=nil {
		fmt.Println("conversion not possible ")
	 }else {fmt.Println(i)}

  // int to string
	 j:=strconv.Itoa(12345)
	 fmt.Println("value of j ", j)
	 fmt.Println("type of j",reflect.TypeOf(j) )


	 //ascii to int
	 char := 'A'
   jl := int(char)
	 fmt.Println("value of datatype : ", jl)


}