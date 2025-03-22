package main

import (
	"fmt"
	"goStructsHandsOn/internal/services"
)

func main() {
	const sep = "-----------------------------------------------------------"
	studentsList := services.NewStudentCollection()

	services.Add(studentsList)

	const a = "Searching..."
	fmt.Println(a)
	s1, err1 := studentsList.SearchStudent("Sagar", "Engineering")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(s1)
	}

	fmt.Println(sep)

	fmt.Println(a)
	s2, err2 := studentsList.SearchStudent("Rohit", "Science")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(s2)
	}

	fmt.Println(sep)

	studentsList.ShowResults()

}
