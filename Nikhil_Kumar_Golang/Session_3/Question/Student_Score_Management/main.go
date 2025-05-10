package main

import (
	"excersices_2/display"
	"excersices_2/student"
	"fmt"
)

func main() {
	for{
	var option int
	fmt.Println(display.Display())
	fmt.Scan(&option)
	if option == 1 {
		student.AddStudent()
	}else if option == 2 {
		student.ViewStudent()
	}else if option == 3 {
		student.UpdateStudentScore()
	}else if option == 4 {
		student.DeleteStudent()
	}else if option == 5 {
		break
	}else {
		fmt.Println("Invalid option")
	}
}
}