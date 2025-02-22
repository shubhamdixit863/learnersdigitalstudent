package main

import (
	"fmt"
	"os"
)

func bankingApp() {
	fmt.Println("enter your name age balance")
	if len(os.Args)> 3{
		fmt.Println("name "+ os.Args[1])
		fmt.Println("age "+ os.Args[2])
		fmt.Println("balance "+ os.Args[3])
		
	}
}



func calc(){
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
		fmt.Printf("addition : %d\n", a+b)
	}else if ch == "-" {
		fmt.Printf("subtraction : %d\n", a-b)
	}else if ch == "*" {
		fmt.Printf("multiplication : %d\n", a*b)
	}else if ch == "/" {
		fmt.Printf("division : %f\n", float64(a) / float64(b))
	}else {	
		fmt.Printf("bye")
	}
}

func fizzbuzz(number int){
	if number % 3 == 0 && number % 5 == 0 {
		fmt.Println("fizzbuzz")
	}else if number % 3 == 0 {
		fmt.Println("Fizz")
	}else if number % 5 == 0 {
		fmt.Println("buzz")
	}else{
		fmt.Println("wrong input")
	}
}

func input() {
	var name string
	var age int
	var address string
	
	fmt.Println("Enter your name:")
	fmt.Scan(&name)

	fmt.Println("Enter your age:")
	fmt.Scan(&age)
	
	fmt.Println()
	var age_in_days int = age * 365
	
	fmt.Print("Enter your address:")
	fmt.Scan(&address)
	
	fmt.Println()
	fmt.Println("**Preson Details**")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Age in days: " , age_in_days)
	fmt.Println("Address:", address)
} 

func main() {
	bankingApp()
	calc()
	fizzbuzz(3)
	fizzbuzz(5)
	fizzbuzz(15)
	fizzbuzz(4)
	input()
}