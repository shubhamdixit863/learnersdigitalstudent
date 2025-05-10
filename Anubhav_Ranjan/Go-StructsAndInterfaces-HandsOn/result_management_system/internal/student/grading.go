package student

import (
	"result_management_system/internal/domain"
)

const (
	MinSubjects         = 1
	PassingGradeDefault = 40
	PassingGradeArts    = 35
)

type Grading interface {
	CalculateGPA(subjects []domain.Subject) (float64, error)
	HasPassed(subjects []domain.Subject) (bool, error)
	GetPassingGrade() int
}

type EngineeringGrading struct{}

func (e *EngineeringGrading) CalculateGPA(subjects []domain.Subject) (float64, error) {
	if len(subjects) < MinSubjects {
		return 0, domain.ErrInvalidSubjects
	}

	var total int
	for _, subject := range subjects {
		total += subject.Grade
	}

	return float64(total) / float64(len(subjects)) / 10, nil
}

func (e *EngineeringGrading) HasPassed(subjects []domain.Subject) (bool, error) {
	if len(subjects) < MinSubjects {
		return false, domain.ErrInvalidSubjects
	}

	for _, subject := range subjects {
		if subject.Grade < PassingGradeDefault {
			return false, nil
		}
	}

	return true, nil
}

func (e *EngineeringGrading) GetPassingGrade() int {
	return PassingGradeDefault
}

type ArtsGrading struct{}

func (a *ArtsGrading) CalculateGPA(subjects []domain.Subject) (float64, error) {
	if len(subjects) < MinSubjects {
		return 0, domain.ErrInvalidSubjects
	}

	var total int
	for _, subject := range subjects {
		total += subject.Grade
	}

	return float64(total) / float64(len(subjects)) / 25, nil
}

func (a *ArtsGrading) HasPassed(subjects []domain.Subject) (bool, error) {
	if len(subjects) < MinSubjects {
		return false, domain.ErrInvalidSubjects
	}

	var total int
	for _, subject := range subjects {
		total += subject.Grade
	}

	average := float64(total) / float64(len(subjects))

	return average >= PassingGradeArts, nil
}

func (a *ArtsGrading) GetPassingGrade() int {
	return PassingGradeArts
}
