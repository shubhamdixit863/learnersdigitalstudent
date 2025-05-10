package student

import (
	"result_management_system/internal/domain"
	"result_management_system/internal/validation"
	"strings"
)

type ArtsStudent struct {
	domain.BaseStudent
	grading Grading
}

func NewArtsStudent(id string, name string) (*ArtsStudent, error) {
	if err := validation.ValidateStudent(id, name); err != nil {
		return nil, err
	}

	return &ArtsStudent{
		BaseStudent: domain.BaseStudent{
			ID:       id,
			Name:     name,
			Subjects: []domain.Subject{},
		},
		grading: &ArtsGrading{},
	}, nil
}

func (a *ArtsStudent) GetID() string {
	return a.ID
}

func (a *ArtsStudent) GetName() string {
	return a.Name
}

func (a *ArtsStudent) GetSubjects() []domain.Subject {
	subjects := make([]domain.Subject, len(a.Subjects))
	copy(subjects, a.Subjects)
	return subjects
}

func (a *ArtsStudent) AddSubject(name string, grade int) error {
	if err := validation.ValidateGrade(grade); err != nil {
		return err
	}

	for _, subject := range a.Subjects {
		if strings.EqualFold(subject.Name, name) {
			return domain.ErrDuplicateSubject
		}
	}

	a.Subjects = append(a.Subjects, domain.Subject{
		Name:  name,
		Grade: grade,
	})

	return nil
}

func (a *ArtsStudent) UpdateGrade(subject string, grade int) error {
	if err := validation.ValidateGrade(grade); err != nil {
		return err
	}

	for i, sub := range a.Subjects {
		if strings.EqualFold(sub.Name, subject) {
			a.Subjects[i].Grade = grade
			return nil
		}
	}

	return domain.ErrSubjectNotFound
}

func (a *ArtsStudent) RemoveSubject(subject string) error {
	for i, sub := range a.Subjects {
		if strings.EqualFold(sub.Name, subject) {
			a.Subjects = append(a.Subjects[:i], a.Subjects[i+1:]...)
			return nil
		}
	}

	return domain.ErrSubjectNotFound
}

func (a *ArtsStudent) CalculateGPA() (float64, error) {
	return a.grading.CalculateGPA(a.Subjects)
}

func (a *ArtsStudent) HasPassed() (bool, error) {
	return a.grading.HasPassed(a.Subjects)
}

func (a *ArtsStudent) GetStudentType() string {
	return "Arts"
}
