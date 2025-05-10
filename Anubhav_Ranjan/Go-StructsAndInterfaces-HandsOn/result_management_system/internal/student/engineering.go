package student

import (
	"result_management_system/internal/domain"
	"result_management_system/internal/validation"
	"strings"
)

type EngineeringStudent struct {
	domain.BaseStudent
	grading Grading
}

func NewEngineeringStudent(id string, name string) (*EngineeringStudent, error) {
	if err := validation.ValidateStudent(id, name); err != nil {
		return nil, err
	}

	return &EngineeringStudent{
		BaseStudent: domain.BaseStudent{
			ID:       id,
			Name:     name,
			Subjects: []domain.Subject{},
		},
		grading: &EngineeringGrading{},
	}, nil
}

func (e *EngineeringStudent) GetID() string {
	return e.ID
}

func (e *EngineeringStudent) GetName() string {
	return e.Name
}

func (e *EngineeringStudent) GetSubjects() []domain.Subject {
	subjects := make([]domain.Subject, len(e.Subjects))
	copy(subjects, e.Subjects)
	return subjects
}

func (e *EngineeringStudent) AddSubject(name string, grade int) error {
	if err := validation.ValidateGrade(grade); err != nil {
		return err
	}

	// Check for duplicate subjects
	for _, subject := range e.Subjects {
		if strings.EqualFold(subject.Name, name) {
			return domain.ErrDuplicateSubject
		}
	}

	e.Subjects = append(e.Subjects, domain.Subject{
		Name:  name,
		Grade: grade,
	})

	return nil
}

func (e *EngineeringStudent) UpdateGrade(subject string, grade int) error {
	if err := validation.ValidateGrade(grade); err != nil {
		return err
	}

	for i, sub := range e.Subjects {
		if strings.EqualFold(sub.Name, subject) {
			e.Subjects[i].Grade = grade
			return nil
		}
	}

	return domain.ErrSubjectNotFound
}

func (e *EngineeringStudent) RemoveSubject(subject string) error {
	for i, sub := range e.Subjects {
		if strings.EqualFold(sub.Name, subject) {
			e.Subjects = append(e.Subjects[:i], e.Subjects[i+1:]...)
			return nil
		}
	}

	return domain.ErrSubjectNotFound
}

func (e *EngineeringStudent) CalculateGPA() (float64, error) {
	return e.grading.CalculateGPA(e.Subjects)
}

func (e *EngineeringStudent) HasPassed() (bool, error) {
	return e.grading.HasPassed(e.Subjects)
}

func (e *EngineeringStudent) GetStudentType() string {
	return "Engineering"
}
