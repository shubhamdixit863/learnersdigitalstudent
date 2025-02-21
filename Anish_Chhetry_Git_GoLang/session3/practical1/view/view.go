package view

import "fmt"

func ViewStudent(names []string, scores []int) {
	fmt.Println("Student Names\tScores")
	for i := 0; i < len(names); i++ {
		fmt.Printf("%d.%s\t\t%d\n", i+1, names[i], scores[i])
	}
}
