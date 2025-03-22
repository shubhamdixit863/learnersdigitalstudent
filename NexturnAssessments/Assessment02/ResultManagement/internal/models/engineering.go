package models

import "fmt"


const (
	EngGradeA = 90.0
	EngGradeB = 80.0
	EngGradeC = 70.0
	EngGradeD = 0.0
)



type EngineeringStudent struct {
	Student
	Electronics float64
	Signals     float64
	Vlsi        float64
}

const some="Engineering Student: %s, ID: %d, Grade: %s"
func (e *EngineeringStudent) CalculateGrade() string {
	average := (e.Electronics+ e.Signals + e.Vlsi) / 3

	switch {
	case average >= EngGradeA:
		return "A"
	case average >= EngGradeB:
		return "B"
	case average >= EngGradeC:
		return "C"
	default:
		return "D"
	}
}

func (e EngineeringStudent) GetDetails() string {
	return fmt.Sprintf(some, e.Name, e.ID, e.CalculateGrade())
}
