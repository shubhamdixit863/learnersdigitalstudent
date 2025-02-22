package main

import "fmt"

func main() {

	var num1 int
	var num2 int
	var oper int

	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	fmt.Println("Enter the operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Scanln(&oper)

	if(oper == 1){
		fmt.Println(num1+num2)
	}else if(oper == 2){
		fmt.Println(num1-num2)
	}else if(oper == 3){
		fmt.Println(num1 * num2)
	}else if(oper == 4){
		if(num2 == 0){
			fmt.Println("Zero division error!")
		}else{
			fmt.Println(num1/num2)
		}
	}else{
		fmt.Println("Enter the valid option!")
	}
}