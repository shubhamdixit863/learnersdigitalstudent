package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Write your Name, age and Bank balance in order after go run main.go.")
	
	if len(os.Args) < 4{
		fmt.Println("Enter the total data!")
	}else{
		fmt.Println("Your name is as follows:")
		fmt.Println(os.Args[1])
		fmt.Println("Your age is as follows:")
		fmt.Println(os.Args[2])
		fmt.Println("Your bank balace is as follows")
		fmt.Println(os.Args[3])
	}
}