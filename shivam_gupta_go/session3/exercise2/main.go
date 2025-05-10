// Student Score Manager:

// 1. Add Student

// 2. View Students

// 3. Update Student Score

// 4. Delete Student

// 5. Exit

package main

import (
	"exercise2/student"
	"fmt"
)

func main() {
	for {
		fmt.Println("1. add student")
		fmt.Println("2. View student")
		fmt.Println("3. update student score")
		fmt.Println("4. delete student ")
		fmt.Println("5. exit")
		var a int
		fmt.Scanln(&a)
		if a == 1 {
			student.Addstudent()
		} else if a == 2 {
			student.ViewStudent()
		} else if a == 3 {
			fmt.Println("enter name ")
			var a string
			fmt.Scanln(&a)

			for i := 0; i < len(student.Names); i++ {
				if a == student.Names[i] {
					fmt.Println("enter score")
					var score int 
					fmt.Scanln(&score)
					student.UpdateScore(i, score)
				}
			}
		} else if a == 4 {
			student.DeleteStudent()
		} else if a == 5 {
			break
		} else {
			fmt.Println("invalid input")
		}
	}

}
