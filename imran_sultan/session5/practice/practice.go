package practice

import (
	"fmt"
	"sort"
)

type Employee struct {
	employe_id     string
	employe_name   string
	employe_salary float32
}

// Define a Student struct with Name, Marks, and Grade. Write a function to assign grades based on marks.
type Student struct {
	name  string
	marks int
	grade string
}

// Create a Rectangle struct with Length and Width. Write a method to calculate the area.
type Rectangle struct {
	length float32
	bredth float32
}

// Create a Book struct with fields Title, Author, and ISBN. Write a function to search for a book by title in a slice of books.
type Book struct {
	Title  string
	Author string
}

// Create a phone directory using a map where names are keys and phone numbers are values. Implement functions to add, delete, and search for contact
func addTomap(mp map[string]string, name string, ph string) map[string]string {
	mp[name] = ph
	return mp
	// fmt.Println(mp)

}
func deletemap(mp map[string]string, name string) map[string]string {
	_, found := mp[name]
	if found {
		delete(mp, name)
	}
	return mp
}

func calArea(l float32, b float32) float32 {
	area := l * b
	return area

}

func setGrade(marks int) string {
	grade := ""
	if marks < 33 {
		grade = grade + "F"
	} else if marks <= 60 {
		grade = grade + "C"
	} else if marks <= 70 {
		grade = grade + "B"
	} else if marks <= 100 {
		grade = grade + "A"
	}
	return grade
}

func Salary_hike(sal int) int {
	if sal < 50000 {
		sal = sal + 10000
	}
	// fmt.Println(sal)
	return sal
}
func checkAuthur(my_book string) string {
	var auth string
	details := []Book{{Title: "insaan e kamil", Author: "murtaza mutheri"}, {Title: "40 hadees", Author: "aga ruhallah komanie"}, {Title: "Zindaan se parey", Author: "syed ali khamenie"}, {Title: "tawazil masiel", Author: "aga sistani"}}
	for i := 0; i < len(details); i++ {
		if details[i].Title == my_book {
			auth = details[i].Author
			break
		}
	}
	return auth
}
func intersection(arr1 []int, arr2 []int) {
	sort.Ints(arr1)
	sort.Ints(arr2)
	i := 0
	j := 0
	k := 0
	l := len(arr1) + len(arr2)
	var new_arr = make([]int, l)
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			new_arr[k] = arr1[i]
			i++
			j++
			k++
		}
	}
	fmt.Println(new_arr)
}
func SetValue() {
	// emp1 := Employee{employe_id: "imr2025", employe_name: "imran", employe_salary: 44000}
	// employee := []Employee{{employe_id: "imr2025", employe_name: "imran", employe_salary: 44000}, {employe_id: "sul2025", employe_name: "sultan", employe_salary: 60000}}
	// Test(emp1)
	// fmt.Println("Before increment")
	// for i := 0; i < len(employee); i++ {
	// 	fmt.Println(employee[i])
	// }
	// fmt.Println("After increment")
	// for i := 0; i < len(employee); i++ {

	// 	newsalary := Salary_hike(int(employee[i].employe_salary))
	// 	employee[i].employe_salary = float32(newsalary)

	// 	fmt.Println(employee[i])
	// }
	// student := []Student{{name: "imran", marks: 98}, {name: "sultan", marks: 60}, {name: "joe", marks: 90}}
	// for i := 0; i < len(student); i++ {
	// 	grd := setGrade(student[i].marks)
	// 	student[i].grade = grd
	// }
	// for i := 0; i < len(student); i++ {
	// 	fmt.Println(student[i])
	// }
	// dimenssions := []Rectangle{{length: 10, bredth: 98}, {length: 23, bredth: 60}, {length: 12, bredth: 90}}
	// for i := 0; i < len(dimenssions); i++ {
	// 	fmt.Println("Area: ", calArea(dimenssions[i].length, dimenssions[i].bredth))
	// }

	// my_books := []string{"insaan e kamil", "40 hadees", "Zindaan se parey", "tawazil masiel"}
	// for i := 0; i < len(my_books); i++ {
	// 	fmt.Println(checkAuthur(my_books[i]))
	// }
	// m1 := make(map[string]string)
	// var choice int
	// for {
	// 	fmt.Println("1 to add	2 to delete 3 to exit")
	// 	fmt.Scanln(&choice)
	// 	switch choice {
	// 	case 1:
	// 		var name string
	// 		fmt.Println("enter name")
	// 		fmt.Scanln(&name)
	// 		var number string
	// 		fmt.Println("Enter Number")
	// 		fmt.Scanln(&number)
	// 		fmt.Println(addTomap(m1, name, number))
	// 	case 2:
	// 		var name string
	// 		fmt.Println("Enter the name you want to delete")
	// 		fmt.Scanln(&name)
	// 		fmt.Println(deletemap(m1, name))
	// 	case 3:
	// 		return
	// 	}
	// }
	arr1 := []int{1, 2, 3, 6, 7}
	arr2 := []int{1, 2, 0, 9}
	intersection(arr1, arr2)

}
