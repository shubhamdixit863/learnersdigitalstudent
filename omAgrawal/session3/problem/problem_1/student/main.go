// package main

// import (
// 	"fmt"
// 	"studentManagement/student"
// )

// func main() {

// 	fmt.Println("Welcome To Student management System")

// 	fmt.Println("--------------------------")

// 	fmt.Println("1. Add Student \n2. View Students \n 3. Update Student Score \n 4. Delete Student \n 5. Exit")

// 	for i := 0; i >= 0; i++ {
// 		fmt.Println("Enter your choice")

// 		var choice int

// 		fmt.Scanln(&choice)

// 		if choice == 1 {
// 			var name string
// 			var score int
// 			fmt.Print("enter Student Name: ")
// 			fmt.Scanln(&name)
// 			fmt.Print("enter Student Score: ")
// 			fmt.Scanln(&score)

// 			student.AddStudent(name, score)
// 		}

// 		if choice == 2 {
// 			student.VeiwAll()
// 		}
// 		if choice == 3 {
// 			fmt.Println("Enter Student name : ")
// 			var studentName string
// 			fmt.Scanln(&studentName)
// 			student.UpdateStudent(studentName)
// 		}
// 		if choice == 4 {
// 			fmt.Println("Enter Student name : ")
// 			var studentName string
// 			fmt.Scanln(&studentName)
// 			student.DeleteStudent(studentName)

// 		}
// 		if choice == 5 {
// 			fmt.Println("Successfully Exited")
// 			return

// 		}
// 	}

// }

package main

import "fmt"

func sliceModifier(a [2]int) {
	// a = append(a, 2)
	a[0] = 9
}

func main() {

	slc := [2]int{2, 3}
	sliceModifier(slc)
	fmt.Println(slc)

}
