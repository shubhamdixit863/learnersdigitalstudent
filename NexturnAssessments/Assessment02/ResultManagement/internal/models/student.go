package models


type Student struct {
	Name   string
	ID     int
	Stream string
}


type Grade interface {
	CalculateGrade() string
	GetDetails() string
}
