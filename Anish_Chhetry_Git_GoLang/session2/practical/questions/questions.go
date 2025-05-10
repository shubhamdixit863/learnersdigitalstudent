package questions

import (
	"assignment2/validate"
	"fmt"
)

var answer int
var correct int

func Question(num int) {
	if num == 1 {
		fmt.Println("What is the capital of France?")
		fmt.Println("1) Berlin")
		fmt.Println("2) Madrid")
		fmt.Println("3) Paris")
		fmt.Println("4) Rome")
		fmt.Printf("Enter your answer (1, 2, 3, or 4): \n")
		correct = 3
		fmt.Scanf("%d", &answer)
		validate.ValidateInput(answer, correct)
	} else if num == 2 {
		fmt.Println("What is 2+2?")
		fmt.Println("1) 5")
		fmt.Println("2) 3")
		fmt.Println("3) 1")
		fmt.Println("4) 4")
		fmt.Printf("Enter your answer (1, 2, 3, or 4): \n")
		correct = 4
		fmt.Scanf("%d", &answer)
		validate.ValidateInput(answer, correct)
	} else if num == 3 {
		fmt.Println("What is the capital of India?")
		fmt.Println("1) New Delhi")
		fmt.Println("2) Mumbai")
		fmt.Println("3) Hyderabad")
		fmt.Println("4) Chennai")
		fmt.Printf("Enter your answer (1, 2, 3, or 4): \n")
		correct = 1
		fmt.Scanf("%d", &answer)
		validate.ValidateInput(answer, correct)
	} else if num == 4 {
		fmt.Println("Choose the odd one out.")
		fmt.Println("1) Dog")
		fmt.Println("2) Cat")
		fmt.Println("3) Shoe")
		fmt.Println("4) Tiger")
		fmt.Printf("Enter your answer (1, 2, 3, or 4): \n")
		correct = 3
		fmt.Scanf("%d", &answer)
		validate.ValidateInput(answer, correct)
	} else if num == 5 {
		fmt.Println("What is 4 divided by 2?")
		fmt.Println("1) 3")
		fmt.Println("2) 2")
		fmt.Println("3) 1")
		fmt.Println("4) 4")
		fmt.Printf("Enter your answer (1, 2, 3, or 4): \n")
		correct = 2
		fmt.Scanf("%d", &answer)
		validate.ValidateInput(answer, correct)
	}
}
