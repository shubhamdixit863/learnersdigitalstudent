package domain

type Student interface {
	GetID() string
	GetName() string
	GetSubjects() []Subject
	AddSubject(name string, grade int) error
	UpdateGrade(subject string, grade int) error
	RemoveSubject(subject string) error
	CalculateGPA() (float64, error)
	HasPassed() (bool, error)
	GetStudentType() string
}

type BaseStudent struct {
	ID       string
	Name     string
	Subjects []Subject
}
