//create a backend application that takes following things:
//first name
//second age
//third balance

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Name:", os.Args[1])
	fmt.Println("Age:", os.Args[2])
	fmt.Println("Balance:", os.Args[3])
}
