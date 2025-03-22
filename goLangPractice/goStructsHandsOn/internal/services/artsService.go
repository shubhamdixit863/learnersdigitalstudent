package services

import "fmt"

type ArtsStudent struct {
	Student
	Score map[string]int
}

func NewArtsStudent(name string, studentId int, score map[string]int) *ArtsStudent {
	return &ArtsStudent{
		Student: Student{
			Name:      name,
			StudentId: studentId,
			Stream:    "Arts",
		},
		Score: score,
	}
}

func (a *ArtsStudent) Grading() string {
	const marksPerSubject = 100
	var marksEarned int
	var totalMarks int
	var percentage float64

	const result = "%s from %s Stream Scored %.2f Percentage"

	for _, score := range a.Score {
		marksEarned += score
		totalMarks += marksPerSubject
	}

	percentage = (float64(marksEarned) / float64(totalMarks)) * 100

	return fmt.Sprintf(result, a.Name, a.Stream, percentage)
}
