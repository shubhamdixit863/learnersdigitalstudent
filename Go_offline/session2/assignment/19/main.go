// Grade Calculator

// Given a score between 0 and 100, assign grades:

// 	•	90-100: A

// 	•	80-89: B

// 	•	70-79: C

// 	•	60-69: D

// 	•	Below 60: F

// Write a function that takes a score and returns the corresponding grade.

package main

import "fmt"

func main() {
	var score int
	fmt.Scanln(&score)
	fmt.Printf("The grade is %s\n", getGrade(score))
}

func getGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}