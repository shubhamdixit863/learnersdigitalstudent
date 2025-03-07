package main

import "fmt"

func main() {

	c := "good"
	fmt.Println(&c)

	c += "morning"
	fmt.Println(&c)

}
