package main

import (
	"fmt"
	"os"
)

func main() {
    fmt.Println("name " + os.Args[1])
		fmt.Println("age " + os.Args[2])
		fmt.Println("balance "+ os.Args[3])
		
}