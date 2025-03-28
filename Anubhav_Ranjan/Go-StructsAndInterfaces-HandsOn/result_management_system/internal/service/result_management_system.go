package service

import (
	"fmt"
	"result_management_system/internal/domain"
	"strings"
)

type ResultManagementSystem struct {
	students map[string]domain.Student
}

func NewResultManagementSystem() *ResultManagementSystem {
	return &ResultManagementSystem{
		students: make(map[string]domain.Student),
	}
}

func (r *ResultManagementSystem) AddStudent(student domain.Student) error {
	if _, exists := r.students[student.GetID()]; exists {
		return fmt.Errorf("student with ID %s already exists", student.GetID())
	}

	r.students[student.GetID()] = student
	return nil
}

func (r *ResultManagementSystem) GetStudent(id string) (domain.Student, error) {
	student, exists := r.students[id]
	if !exists {
		return nil, fmt.Errorf("student with ID %s not found", id)
	}

	return student, nil
}

func (r *ResultManagementSystem) RemoveStudent(id string) error {
	if _, exists := r.students[id]; !exists {
		return fmt.Errorf("student with ID %s not found", id)
	}

	delete(r.students, id)
	return nil
}

func (r *ResultManagementSystem) GetStudentResults(id string) (string, error) {
	student, err := r.GetStudent(id)
	if err != nil {
		return "", err
	}

	gpa, err := student.CalculateGPA()
	if err != nil {
		return "", fmt.Errorf("error calculating GPA: %w", err)
	}

	passed, err := student.HasPassed()
	if err != nil {
		return "", fmt.Errorf("error determining pass status: %w", err)
	}

	var status string
	if passed {
		status = "PASSED"
	} else {
		status = "FAILED"
	}

	var result strings.Builder
	result.WriteString(fmt.Sprintf("Student ID: %s\n", student.GetID()))
	result.WriteString(fmt.Sprintf("Name: %s\n", student.GetName()))
	result.WriteString(fmt.Sprintf("Type: %s\n", student.GetStudentType()))
	result.WriteString(fmt.Sprintf("GPA: %.2f\n", gpa))
	result.WriteString(fmt.Sprintf("Status: %s\n", status))
	result.WriteString("Subjects:\n")

	for _, subject := range student.GetSubjects() {
		result.WriteString(fmt.Sprintf("- %s: %d\n", subject.Name, subject.Grade))
	}

	return result.String(), nil
}
