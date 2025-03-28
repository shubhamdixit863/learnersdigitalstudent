package main

import (
	"fmt"
	"module_Assignment/exit"
	"module_Assignment/quize"
	"module_Assignment/score"
)

func main() {

	fmt.Println(" ----- Quize Game------------------------")
	// var input int

	var input int
	// score.Score


	for i := 0; i >= 0; i++ {

		fmt.Println("1. Start quize")
		fmt.Println("2. veiw Score")
		fmt.Println("3. Exit")
		fmt.Println("enter your preference")
		fmt.Scanln(&input)

		if input == 1 {
			quize.Quize()

		} else if input == 2 {
			score.ScorePrint()
		} else if input == 3 {
			exit.Exit()
			return
		}
	}

}
