package main 
import "fmt"

func main(){

	fmt.Println("select operations")
	var op string
	fmt.Scan(&op)
	fmt.Println("enter two numbers")
	var a float32 
	var b float32
	fmt.Scan(&a)
	fmt.Scan(&b)
	if op=="+" {
		fmt.Println(" ans", a+b)
	}else if op =="-" {fmt.Println("ans",a-b)}else if op == "*"{
		fmt.Println("ans", a*b)
	} else if op =="/"{
		fmt.Println("ans", a/b)
	}else {
		fmt.Println("error")
	}


}