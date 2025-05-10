package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) float32 {
	return float32(a)/float32(b)
}


func main() {
	var a, b int
	fmt.Print("Enter two numbers ")
	fmt.Print("A : ")
	fmt.Scan(&a)
	fmt.Print("B : ")
	fmt.Scan(&b)
	fmt.Print(" + for addition,\n - for subtraction \n * for multiplication \n / for division \n enter your operation : ")
	var ch string
	fmt.Scan(&ch)
	if ch == "+" {
		fmt.Printf("addition : %d\n", add(a, b))
	}else if ch == "-" {
		fmt.Printf("subtraction : %d\n", sub(a, b))
	}else if ch == "*" {
		fmt.Printf("multiplication : %d\n", mul(a, b))
	}else if ch == "/" {
		fmt.Printf("division : %f\n", div(a, b))
	}else {	
		fmt.Printf("bye")
	}
}