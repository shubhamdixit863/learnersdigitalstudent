package main

import "fmt"

var stud_name []string
var stud_marks []int

// var options []string={
// 		"choise 1 to add"
// 		"choise 2 to view"
// 		"choise 3 to update"
// 		"choise 4 to delete"
// 		"choise 5 to exit"
// }

func add_student() {
	var name string
	var marks int
	fmt.Println("Enter the your name and marks")
	fmt.Scanln(&name)
	fmt.Scanln(&marks)
	stud_name = append(stud_name, name)
	stud_marks = append(stud_marks, marks)

	// return stud_name, stud_marks
}
func viewAll() {
	for i := 0; i < len(stud_name); i++ {
		fmt.Println(stud_name[i], stud_marks[i])
	}
}
func update() {
	fmt.Println("Enter the name of student you want to update: ")
	var name string
	fmt.Scanln(&name)
	for i := 0; i < len(stud_name); i++ {
		if stud_name[i] == name {
			fmt.Println("Enter the new marks of ", name)
			var new_marks int
			fmt.Scanln(&new_marks)
			stud_marks[i] = new_marks
			return
		}
	}
}
func delete() {
	fmt.Println("Enter the name of student you want to delete: ")
	var name string
	fmt.Scanln(&name)
	for i := 0; i < len(stud_name); i++ {
		if stud_name[i] == name {
			stud_name = append(stud_name[:i], stud_name[i+1:]...)
			stud_marks = append(stud_marks[:i], stud_marks[i+1:]...)
			return
		}
	}
}
func upd(name []string) []string {
	name = append(name, "imran khan")
	return name
}

func main() {

	// for {
	// 	println("choise 1 to add")
	// 	println("choise 2 to view")
	// 	println("choise 3 to update")
	// 	println("choise 4 to delete")
	// 	println("choise 5 to exit")
	// 	var want int
	// 	fmt.Scanln(&want)
	// 	switch want {
	// 	case 1:
	// 		add_student()
	// 	case 2:
	// 		viewAll()
	// 	case 3:
	// 		update()
	// 	case 4:
	// 		delete()
	// 	case 5:
	// 		return
	// 	}

	// }
	name := []string{"imran"}

	upd(name)
	fmt.Println(name)

}
