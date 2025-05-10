// Add a Student:

// 	•	Prompt the user to enter a student name and score.

// 	•	Store names in a []string slice and scores in a []int slice.

// 	•	Names and scores should be aligned by index (e.g., names[0] corresponds to scores[0])

package student

import (
	
	"fmt"
)

var Names []string
var Score []int

func Addstudent() []string {
	fmt.Println("enter student name : ")
	var a string
	var b int
	fmt.Scanln(&a)
	fmt.Println("enter score ")
	fmt.Scanln(&b)
	Names = append(Names, a)
	Score = append(Score, b)
	fmt.Println("student added successfully")
	return Names
}

func DeleteStudent() {
	fmt.Println("enter student name to delete")
	var name string
	fmt.Scanln(&name)
	var i int
	for i = 0; i < len(Names); i++ {
		if Names[i] == name {
			break
		}
	}
	Names = append(Names[:i], Names[i+1:]...)
	Score = append(Score[:i], Score[i+1:]...)
	fmt.Println("deleted successfully")
}

func ViewStudent() {
	if len(Names) == 0 {
		fmt.Println("no st")
	}
	for i := 0; i < len(Names); i++ {

		fmt.Println(Names[i], Score[i])
	}
}

func UpdateScore(index int , val int ){
  
	Score[index]=val

	fmt.Println("update successfully")
}