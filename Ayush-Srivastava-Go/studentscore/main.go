package main

import (
	"fmt"
	"studentscore/constants"
	"studentscore/students"
)

func main() {

	for {
		var choice int

		fmt.Println("\nStudent Score Manager:")
		
		for idx := range constants.Choices{
			fmt.Println(constants.Choices[idx])
		}

		fmt.Println("\nEnter your choice")
		fmt.Scan(&choice)

		switch choice{
		case 1:
			students.AddStudent()
		case 2:
			students.ViewStudents()
		case 3: 
			students.UpdateScore()
		case 4: 
			students.DeleteStudent()
		case 5:
			fmt.Println("See yaa later!")
			return
		default:
			fmt.Println("Please enter a valid choice!")
			return
		}
	}

}