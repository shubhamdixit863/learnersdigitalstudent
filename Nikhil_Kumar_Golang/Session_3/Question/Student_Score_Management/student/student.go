package student

import (
	"excersices_2/score"
	"fmt"
)

var Student []string

func ViewStudent() {
	if len(Student) == 0 {
		fmt.Println("No student found")
	}
	for i := 0; i < len(Student); i++ {
		fmt.Println(i+1,". ", Student[i]," - ",score.Score[i])
	}
}

func AddStudent() {
	var name string
	var sc int
	ln := len(Student)
	fmt.Println("Enter Student name :")
	fmt.Scan(&name)
	Student = append(Student, name)
	fmt.Println("Enter Score : ")
	fmt.Scan(&sc)
	score.AddScore(sc)
	ln2 := len(Student)
	if ln == (ln2-1) {
		fmt.Println("Student Added Successfully")
	}
}

func UpdateStudentScore() {
	var name string
	fmt.Println("Enter Student Name to Update : ")
	fmt.Scan(&name)

	for i := 0; i < len(Student); i++ {
		if Student[i] == name {
			fmt.Println("Enter New Score : ")
			var sc int
			fmt.Scan(&sc)
			score.UpdatedScore(sc, i)
		}else {
			fmt.Println("Student not found")
			break
		}
	}
	fmt.Println("Score updated successfully!")

}

func DeleteStudent() {
	var name string
	fmt.Println("Enter Student Name to Update : ")
	fmt.Scan(&name)

	for i := 0; i < len(Student); i++ {
		if Student[i] == name {
			Student = append(Student[:i], Student[i+1:]...)
			score.Score = append(score.Score[:i], score.Score[i+1:]...)
		}
}
fmt.Println(" Student Deleted Successfully")


}