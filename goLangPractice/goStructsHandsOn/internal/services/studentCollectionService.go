package services

import (
	"errors"
	"fmt"
)

type StudentCollection struct {
	EngStudents  []EngineeringStudent
	ArtsStudents []ArtsStudent
}

func NewStudentCollection() *StudentCollection {
	return &StudentCollection{}
}

func (s *StudentCollection) AddStudent(student interface{}) {
	const result = "Added Student: %v\n"
	const r = "Student %v has no implementation\n"

	switch std := student.(type) {
	case *EngineeringStudent:
		s.EngStudents = append(s.EngStudents, *std)
		fmt.Printf(result, *std)
	case *ArtsStudent:
		s.ArtsStudents = append(s.ArtsStudents, *std)
		fmt.Printf(result, *std)
	default:
		fmt.Printf(r, std)
	}

}

func (s *StudentCollection) ShowResults() {
	const x = "The Following are the results of engineering students: "
	const y = "The Following are the results of arts students: "
	const sep = "-----------------------------------------------------------"

	fmt.Println(x)
	for _, e := range s.EngStudents {
		fmt.Println(e.Grading())
	}

	fmt.Println(sep)

	fmt.Println(y)
	for _, a := range s.ArtsStudents {
		fmt.Println(a.Grading())
	}
}

func (s *StudentCollection) SearchStudent(name string, stream string) (interface{}, error) {
	const err = "student not found"
	const e = "not a valid stream"
	if stream == "Engineering" {
		for _, e := range s.EngStudents {
			if e.Name == name {
				return e, nil
			} else {
				return nil, errors.New(err)
			}
		}
	} else if stream == "Arts" {
		for _, a := range s.ArtsStudents {
			if a.Name == name {
				return a, nil
			} else {
				return nil, errors.New(err)
			}
		}
	}
	return nil, errors.New(e)
}
