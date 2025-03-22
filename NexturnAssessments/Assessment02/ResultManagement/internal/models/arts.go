package models

import "fmt"

const(
	 ArtsGradeA =85.0
	 ArtsGradeB =75.0
	 ArtsGradeC = 65.0
	 ArtsGradeD =0.0

)


type ArtsStudent struct {
	Student
	History   float64
	Geography float64
	Civics    float64
}


func (a *ArtsStudent) CalculateGrade() string {
	average := (a.History + a.Geography+ a.Civics) / 3
	switch {
	case average >= ArtsGradeA:
		return "A"
	case average >= ArtsGradeB:
		return "B"
	case average >= ArtsGradeC:
		return "C"
	default:
		return "D"
}
}

func (a ArtsStudent) GetDetails() string {
	return fmt.Sprintf("Arts Student: %s, ID: %d, Grade: %s", a.Name, a.ID, a.CalculateGrade())
}
