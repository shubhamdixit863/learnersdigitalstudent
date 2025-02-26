package student_grade_management_system

import (
	"errors"
	"fmt"
)

type Student struct {
	Name  string
	Marks float64
	Grade rune
}

func AssignGrade(marks float64) rune {
	switch {
	case marks >= 90:
		return 'A'
	case marks >= 75:
		return 'B'
	case marks >= 60:
		return 'C'
	case marks >= 35:
		return 'D'
	default:
		return 'F'
	}
}

func AddStudent(db *[]Student, name string, marks float64) {
	grade := AssignGrade(marks)
	newStudent := Student{Name: name, Marks: marks, Grade: grade}
	*db = append(*db, newStudent)
	fmt.Printf(" Student Added: %+v\n", newStudent)
}

func RemoveStudent(db *[]Student, name string) error {
	for i, stud := range *db {
		if stud.Name == name {
			*db = append((*db)[:i], (*db)[i+1:]...)
			fmt.Printf(" Student with Name %s is Removed..\n", name)
			return nil
		}
	}

	return errors.New("Student NOT Found")
}

func DisplayStudents(db *[]Student) {
	if len(*db) == 0 {
		fmt.Println("No Student Found")
		return
	}

	fmt.Println("\n Student Records:")
	for _, stud := range *db {
		fmt.Printf("Name : %s | Marks : %.2f | Grade : %c\n", stud.Name, stud.Marks, stud.Grade)
	}
}
