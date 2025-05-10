package main

import (
	"fmt"
	"studentscore/student"
)

func main() {
	for {
		fmt.Println("\nQuiz Game")
		fmt.Println("1. Add Student")
		fmt.Println("2. Add Score")
		fmt.Println("3. Update Student Score")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var p int
			fmt.Println("Enter the number of Students")
			fmt.Scanln(&p)
			student.Addstudent(p)

		case 2:
			var p int
			fmt.Println("Enter the score of Students")
			fmt.Scanln(&p)
			student.Viewstudent(p)

		case 3:
			var name string
			fmt.Println("Enter name of student to update")
			fmt.Scanln(&name)
			student.Updatescore(name)

		case 4:
			var name string
			fmt.Println("Enter name of student to delete")
			fmt.Scanln(&name)
			student.Deletestudent(name)

		case 5:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice, try again!")
		}
	}
}
