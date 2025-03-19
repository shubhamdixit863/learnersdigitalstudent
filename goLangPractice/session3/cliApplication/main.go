package main

import "fmt"

var Name = []string{}
var Score = []int{}

func addStudent() {
	var name string
	var score int

	fmt.Print("Enter Student Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Score: ")
	fmt.Scan(&score)

	Name = append(Name, name)
	Score = append(Score, score)
	fmt.Println("Student Added Successfully")
}

func viewStudents() {
	fmt.Println("Students List With Scores: ")

	if len(Name) == 0 {
		fmt.Println("No Students Found")
		return
	}

	for i, _ := range Name {
		fmt.Printf("%s - %d\n", Name[i], Score[i])
	}
}

func updateScore() {
	var name string
	var score int

	fmt.Print("Enter Student Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter New Score: ")
	fmt.Scan(&score)

	for i, _ := range Name {
		if Name[i] == name {
			Score[i] = score
			fmt.Println("Student Score Updated Successfully")
			return
		}
	}
	fmt.Println("Updation Failed: No Students Found With Name: ", name)
}

func deleteStudent() {
	var name string
	fmt.Print("Enter Student Name: ")
	fmt.Scan(&name)

	for i, _ := range Name {
		if Name[i] == name {
			Name = append(Name[:i], Name[i+1:]...)
			Score = append(Score[:i], Score[i+1:]...)
			fmt.Println("Student Deleted Successfully")
			return
		}
	}
	fmt.Println("Deletion Failed: No Students Found With Name: ", name)
}

func main() {
	temp := true
	fmt.Println("Student Score Manager:\n\n1. Add Student\n\n2. View Students\n\n3. Update Student Score\n\n4. Delete Student\n\n5. Exit")

	for temp {
		var choice int
		fmt.Print("Enter Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			viewStudents()
		case 3:
			updateScore()
		case 4:
			deleteStudent()
		case 5:
			fmt.Println("Exit")
			temp = false
		default:
			fmt.Println("Invalid Choice")

		}

	}
}
