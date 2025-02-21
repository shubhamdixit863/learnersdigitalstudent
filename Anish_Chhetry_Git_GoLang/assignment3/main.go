package main

import (
	"assignment2/score"
	"assignment2/startquiz"
	"fmt"
)

func main() {
	choose := 0
	for i := 0; i == 0; {
		fmt.Printf("=== Quiz Game ===\n1. Start Quiz\n2. View Score\n3. Exit\nEnter your choice:)")
		fmt.Scanf("%d", &choose)
		switch choose {
		case 1:
			score.ResetScore()
			startquiz.Startquiz()
		case 2:
			score.ViewScore()
		case 3:
			i = 1
			fmt.Println("Thank you for playing the Quiz Game.")
		}
	}
}
