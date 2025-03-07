package main

import (
	"exercise2/student"
	"fmt"
)

var options = []string{
	"1. Add Student",
	"2. View Students",
	"3. Update Students",
	"4. Delete Student",
	"5. Exit"}

func main() {

	fmt.Println("\t\tStudent Score Manager:")
	for {
		for i := 0; i < len(options); i++ {
			fmt.Println(options[i])
		}
		fmt.Printf("\n")
		fmt.Printf("Enter Choice: ")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("\n\n")
		switch choice {
		case 1:
			student.AddStudent()
			fmt.Printf("\n\n")
		case 2:
			student.ViewStudent()
			fmt.Printf("\n\n")
		case 3:
			student.UpdateStudent()
			fmt.Printf("\n\n")
		case 4:
			student.DeleteStudent()
			fmt.Printf("\n\n")
		case 5:
			return
		default:
			fmt.Println("Please Enter a valid choice!!")
			fmt.Printf("\n\n")
		}

	}
}
