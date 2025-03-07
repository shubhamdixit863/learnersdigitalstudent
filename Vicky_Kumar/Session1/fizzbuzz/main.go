package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	num, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Fizzbuzz(num))
	}
}
