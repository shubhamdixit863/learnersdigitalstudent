package student

import "fmt"

var names []string
var scores []int

func AddStudent() {
	fmt.Printf("Enter student Name: ")
	var name string
	fmt.Scanln(&name)
	names = append(names, name)
	fmt.Printf("Enter Score: ")
	var score int
	fmt.Scanln(&score)
	scores = append(scores, score)
	fmt.Println("Student added successfully!")
}
func ViewStudent() {
	if len(names) == 0 {
		fmt.Println("No Students Found")
		return
	}
	fmt.Println("Students and Scores:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("%d. %s-%d\n", i+1, names[i], scores[i])
	}
}
func UpdateStudent() {
	fmt.Printf("Enter Student Name to Update:")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Enter New Score:")
	var score int
	fmt.Scanln(&score)
	flag := true
	for i := 0; i < len(names); i++ {
		if names[i] == name {
			scores[i] = score
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("No Student found..!")
		return
	}
	fmt.Println("Score updated successfully!")
}
func DeleteStudent() {
	fmt.Printf("Enter Student Name to Delete:")
	var name string
	fmt.Scanln(&name)
	flag := true
	for i := 0; i < len(names); i++ {
		if names[i] == name {
			names = append(names[:i], names[i+1:]...)
			scores = append(scores[:i], scores[i+1:]...)
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("No Student found..!")
		return
	}
	fmt.Println("Student Deleted successfully!")
}
