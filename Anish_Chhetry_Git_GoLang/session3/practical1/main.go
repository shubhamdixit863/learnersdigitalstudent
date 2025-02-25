package main

import (
	"assignment5/add"
	"assignment5/delete"
	"assignment5/update"
	"assignment5/view"
	"fmt"
)

var names []string
var scores []int

func main() {
	// var names []string
	// var scores []int
	var option int
	for i := 0; i == 0; {
		fmt.Println("\n-----------------------------------------------------------------")
		fmt.Println("Welcome to Student Score Management System!")
		fmt.Printf("Choose one of the following: \n1.Add\n2.View\n3.Update\n4.Delete\n5.Exit\n")
		fmt.Printf("Enter your Choice: ")
		fmt.Scanf("%d\n", &option)
		switch option {
		case 1:
			fmt.Println("\n-----------------------------------------------------------------")
			names, scores = add.AddStudent(names, scores)
			fmt.Println("\n-----------------------------------------------------------------")
		case 2:
			fmt.Println("\n-----------------------------------------------------------------")
			view.ViewStudent(names, scores)
			fmt.Println("\n-----------------------------------------------------------------")
		case 3:
			fmt.Println("\n-----------------------------------------------------------------")
			names, scores = update.UpdateStudent(names, scores)
			fmt.Println("\n-----------------------------------------------------------------")
		case 4:
			fmt.Println("\n-----------------------------------------------------------------")
			names, scores = delete.DeleteStudent(names, scores)
			fmt.Println("\n-----------------------------------------------------------------")
		case 5:
			fmt.Println("Goodbye!")
			i++
		default:
			fmt.Println("Invalid option. Please choose again.")
		}

	}
}
