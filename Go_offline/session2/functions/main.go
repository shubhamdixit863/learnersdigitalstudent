package main

import "fmt"

func main(){
c:=square(3)
d:=sum(3,4)
e, err:=divide(3,4)

// e,_:=divide(3,0) //if you want to ignore the error

//error handling by putting error check in the return type   very important
// fmt.Println(e, err)

if err!=nil {
	fmt.Println(err)
	return
} 
fmt.Println(e)

fmt.Println(c)
fmt.Println(d)
variadicFun(3.4, 1,2,3,4,5,6,7,8,9,10)

f,g:=namedReturn(3,4)
fmt.Println(f,g)
}

func square(a int) int {
	c:=a^2
	return c
}

func sum(a int, b int) int {
	c:=a+b
	return c
}

func divide(a int, b int) (int, error) {
	if b==0 {
		return 0,  fmt.Errorf("cannot divide by zero %s", "error")
	}
	c:=a/b
	return c, nil
}    



func variadicFun(f float64,n ...int ){
	fmt.Println(f, n)
}


func namedReturn(a int, b int) (c int, d int) {  //named return its easier to understand the return values
	c=a+b
	d=a-b
	return
}