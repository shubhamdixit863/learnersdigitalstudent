package student

import (
	"fmt"
)

var students []string
var scores []int

func Addstudent(n int) {
	fmt.Println("Enter", n, "student names:")
	students = make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&students[i])
	}
	fmt.Println("Student names:", students)
}

func Viewstudent(m int) {
	fmt.Println("Enter", m, "student scores:")
	scores = make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&scores[i])
	}
	fmt.Println("Student Scores:", scores)
}

func Updatescore(name string) {
	for i := 0; i < len(students); i++ {
		if name == students[i] {
			var p int
			fmt.Println("Enter new score")
			fmt.Scanln(&p)
			scores[i] = p
			fmt.Println("Updated")
			return
		}
	}
	fmt.Println("Student not found")
}

func Deletestudent(name string) {
	for i := 0; i < len(students); i++ {
		if students[i] == name {
			students = append(students[:i], students[i+1:]...)
			scores = append(scores[:i], scores[i+1:]...)
			fmt.Println("Deleted Student:", name)
			return
		}
	}
	fmt.Println("Student not found")
}
