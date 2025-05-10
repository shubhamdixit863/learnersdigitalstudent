package main

import (
	"fmt"

	"result_management_system/internal/service"
	"result_management_system/internal/student"
)

func main() {
	rms := service.NewResultManagementSystem()

	engineeringStudent, err := student.NewEngineeringStudent("E001", "Anubhav Ranjan")
	if err != nil {
		fmt.Printf("Error creating engineering student: %v\n", err)
		return
	}

	artsStudent, err := student.NewArtsStudent("A001", "Shirish Manoj Bobde")
	if err != nil {
		fmt.Printf("Error creating arts student: %v\n", err)
		return
	}

	if err := rms.AddStudent(engineeringStudent); err != nil {
		fmt.Printf("Error adding engineering student: %v\n", err)
		return
	}

	if err := rms.AddStudent(artsStudent); err != nil {
		fmt.Printf("Error adding arts student: %v\n", err)
		return
	}

	if err := engineeringStudent.AddSubject("Mathematics", 85); err != nil {
		fmt.Printf("Error adding subject: %v\n", err)
	}

	if err := engineeringStudent.AddSubject("Physics", 78); err != nil {
		fmt.Printf("Error adding subject: %v\n", err)
	}

	if err := engineeringStudent.AddSubject("Computer Science", 92); err != nil {
		fmt.Printf("Error adding subject: %v\n", err)
	}

	if err := artsStudent.AddSubject("History", 88); err != nil {
		fmt.Printf("Error adding subject: %v\n", err)
	}

	if err := artsStudent.AddSubject("Literature", 92); err != nil {
		fmt.Printf("Error adding subject: %v\n", err)
	}

	if err := artsStudent.AddSubject("Philosophy", 75); err != nil {
		fmt.Printf("Error adding subject: %v\n", err)
	}

	engineeringResults, err := rms.GetStudentResults("E001")
	if err != nil {
		fmt.Printf("Error getting engineering student results: %v\n", err)
	} else {
		fmt.Println("Engineering Student Results:")
		fmt.Println(engineeringResults)
	}

	artsResults, err := rms.GetStudentResults("A001")
	if err != nil {
		fmt.Printf("Error getting arts student results: %v\n", err)
	} else {
		fmt.Println("Arts Student Results:")
		fmt.Println(artsResults)
	}
}
