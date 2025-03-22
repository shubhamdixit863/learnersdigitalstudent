package main

import (
	"fmt"
	"question/score"
	"question/start"
)

func main() {
	fmt.Println("Enter what you want: ")
	fmt.Println("1-->start")
	fmt.Println("2-->score")
	fmt.Println("3-->exit")
	var choice string
	fmt.Scan(&choice)
	if choice == "1" {
		start.Hello()

	} else if choice == "2" {
		score.ShowScore()
	} else {
		fmt.Println("exit")
	}

}
