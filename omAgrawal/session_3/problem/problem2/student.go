package student

import "fmt"

var Student []string
var Score []int

func AddStudent(name string, score int) {
	Student = append(Student, name)
	Score = append(Score, score)

	fmt.Println("Student added succesfully")

}

func VeiwAll() {
	fmt.Println("no.   Name    Score")
	for i := 0; i < len(Student); i++ {
		fmt.Println(i+1, "   ", Student[i], "   ", Score[i])
	}
}

func UpdateStudent(name string) {
	for i := 0; i < len(Student); i++ {
		if name == Student[i] {
			var score int
			fmt.Println("Enter Score :")
			fmt.Scanln(&score)
			Score[i] = score
			fmt.Println("Student updated succesfully")
			return
		}
	}
	fmt.Println("No Student Found")
}

func DeleteStudent(name string) {
	xer := -1

	for i := 0; i < len(Student); i++ {
		if name == Student[i] {
			xer = i
		}
	}

	Student = append(Student[:xer], Student[xer+1:]...)

	Score = append(Score[:xer], Score[xer+1:]...)
}
