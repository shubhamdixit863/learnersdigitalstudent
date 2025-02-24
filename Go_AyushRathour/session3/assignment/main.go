package main

import (
	"assignment/studentManagement"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Student Score Manager")



	for {

		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Update Student Score")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")

		fmt.Println("Enter your choice:")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			studentManagement.AddStudent()
		case 2:
			studentManagement.ViewStudents()
		case 3:
			studentManagement.UpdateStudentScore()
		case 4:
			studentManagement.DeleteStudent()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}
