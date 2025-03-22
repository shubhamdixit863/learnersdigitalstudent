package services

import "fmt"

func Add(studentsList *StudentCollection) {
	const sep = "-----------------------------------------------------------"

	m1 := map[string]map[int]int{
		"Maths":     {4: 9},
		"Physics":   {5: 8},
		"Chemistry": {3: 7},
	}
	e1 := NewEngineeringStudent("Sagar", 1, m1)
	studentsList.AddStudent(e1)

	fmt.Println(sep)

	m2 := map[string]map[int]int{
		"Maths":     {4: 7},
		"Physics":   {5: 6},
		"Chemistry": {3: 9},
	}
	e2 := NewEngineeringStudent("Rohit", 2, m2)
	studentsList.AddStudent(e2)

	fmt.Println(sep)

	m3 := map[string]int{
		"English": 80,
		"History": 94,
		"Civics":  87,
	}
	a1 := NewArtsStudent("Sahil", 1, m3)
	studentsList.AddStudent(a1)

	fmt.Println(sep)

	m4 := map[string]int{
		"English": 92,
		"History": 77,
		"Civics":  67,
	}
	a2 := NewArtsStudent("Vikas", 2, m4)
	studentsList.AddStudent(a2)

	fmt.Println(sep)
}
