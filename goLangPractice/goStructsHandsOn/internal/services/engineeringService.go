package services

import "fmt"

type EngineeringStudent struct {
	Student
	Score map[string]map[int]int
}

func NewEngineeringStudent(name string, studentId int, score map[string]map[int]int) *EngineeringStudent {
	return &EngineeringStudent{
		Student: Student{
			Name:      name,
			StudentId: studentId,
			Stream:    "Engineering",
		},
		Score: score,
	}
}

func (e *EngineeringStudent) Grading() string {
	var creditsEarned int
	var totalCredits int
	var cgpa float64

	const result = "%s from %s Stream Scored %.2f CGPA out of 10"

	for _, score := range e.Score {
		for credit, pointer := range score {
			totalCredits += credit
			creditsEarned += credit * pointer
		}
	}

	cgpa = float64(creditsEarned) / float64(totalCredits)

	return fmt.Sprintf(result, e.Name, e.Stream, cgpa)
}
