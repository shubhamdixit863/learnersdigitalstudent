package studentManagement

import (
	"fmt"
)

var name []string
var score []int

func AddStudent() {
	var studentName string
	var studentScore int
	fmt.Println("Enter Student :")
	fmt.Scanln(&studentName)
	fmt.Println("Enter Score :")
	fmt.Scanln(&studentScore)
	name=append(name,studentName)
	score=append(score,studentScore)
    fmt.Println("Student Added Successfully!")
}

func ViewStudents() {
	for i:=0;i<len(name);i++{
		fmt.Printf("%s - %d \n", name[i],score[i])
	}
}

func UpdateStudentScore() {
	var studentName string
	var studentScore int
	fmt.Println("Enter Student :")
	fmt.Scanln(&studentName)
	fmt.Println("Enter Score :")
	fmt.Scanln(&studentScore)
	for i:=0;i<len(name);i++{
		if name[i]==studentName{
			score[i]=studentScore
			fmt.Println("Score Updated Successfully!")
			return
		}
	}
	fmt.Println("Student not found!")
}

func DeleteStudent() {
	var studentName string
	fmt.Println("Enter Student :")
	fmt.Scanln(&studentName)
	for i:=0;i<len(name);i++{
		if name[i]==studentName{
			name=append(name[:i],name[i+1:]...)
			score=append(score[:i],score[i+1:]...)
			fmt.Println("Student Deleted Successfully!")
			return
		}
	}
	fmt.Println("Student not found!")
}




