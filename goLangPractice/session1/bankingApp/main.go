package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 3 {
		fmt.Println("Your name is: " + os.Args[1])
		fmt.Println("Your age is: " + os.Args[2])
		fmt.Println("Your balance is: " + os.Args[3])
	}
}
