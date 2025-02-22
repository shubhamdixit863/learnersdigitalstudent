package students

import (
	"fmt"
	"strings"
)

var names []string
var scores []int

func AddStudent() {

	var name string
	var score int

	fmt.Println("Enter student's name: ")
	fmt.Scan(&name)
	
	fmt.Println("Enter student's score: ")
	fmt.Scan(&score)

	name = strings.TrimSpace(name)

	names = append(names, name)
	scores = append(scores, score)
	fmt.Println("Student Saved Successfully !!")
}

func ViewStudents(){
	if(len(names) == 0){
		fmt.Println("No student found!!")
		return
	}

	fmt.Println("Students and Scores: ")
	for i, name := range names{
		fmt.Printf("%d. %s got %d score\n", i+1, name, scores[i])
	}
}

func UpdateScore(){

	var name string
	var score int

	fmt.Println("Enter the name of the student to be edited: ")
	fmt.Scan(&name)

	name = strings.TrimSpace(name)

	for i, n := range names{
		if n == name{
			fmt.Println("Enter the updated score: ")
			fmt.Scan(&score)
			scores[i] = score

			fmt.Println("Score updated successfully!")
			return
		}
	}

	fmt.Println("Student not found!")
}

func DeleteStudent(){

	var name string

	fmt.Println("Enter the name of the student to be deleted: ")
	fmt.Scan(&name)

	name = strings.TrimSpace(name)

	for i, n := range names{
		if n == name{
			names = append(names[:i], names[i+1:]...)
			scores = append(scores[:i], scores[i+1:]...)
			fmt.Print("Student record deleted Successfully!")
			return
		}
	}
	fmt.Println("No such student found!")
}

