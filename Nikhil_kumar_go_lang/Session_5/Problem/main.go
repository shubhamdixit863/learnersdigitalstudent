package main

import (
	"fmt"
	"math"
	"strconv"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func employee_records() {
	emp := Employee{
		ID:     1,
		Name:   "John Doe",
		Salary: 50000.0,
	}
	hike_salary(&emp, 5000.0)
	fmt.Printf("ID: %d\nName: %s\nSalary: %.2f\n", emp.ID, emp.Name, emp.Salary)
}
func hike_salary(emp *Employee, hike float64) {
	emp.Salary += hike
}

type Student struct {
	Name  string
	Marks int
	Grade string
}

func student_grade_management() {
	student := Student{
		Name:  "John Doe",
		Marks: 85,
	}
	fmt.Printf("Name: %s\nMarks: %d\nGrade: %s\n", student.Name, student.Marks, assign_grage(student.Marks))
}

func assign_grage(marks int) string {
	if marks >= 90 {
		return "A"
	} else if marks >= 80 {
		return "B"
	} else if marks >= 70 {
		return "C"
	} else if marks >= 60 {
		return "D"
	} else {
		return "F"
	}
}

type Rectangle struct {
	Length float64
	Width  float64
}

func area_of_rectangle() {
	rectangle := Rectangle{
		Length: 5.0,
		Width:  3.0,
	}
	fmt.Printf("Area of Rectangle: %.2f\n", rectangle.Length*rectangle.Width)
}

type Circle struct {
	Radius float32
}

func circle_perimeter_and_area() {
	circle := Circle{
		Radius: 5.0,
	}
	fmt.Printf("Perimeter of Circle: %.2f\n", 2*math.Pi*circle.Radius)
	fmt.Printf("Area of Circle: %.2f\n", math.Pi*circle.Radius*circle.Radius)
}

type Book struct {
	Title  string
	Author string
	ISBN   int
}

func library_book_management() {
	book := Book{
		Title:  "The Go Programing Language",
		Author: "Alan A. A. Donovan",
		ISBN:   1234567890123,
	}
	fmt.Printf("Title: %s\nAuthor: %s\nISBN: %d\n", book.Title, book.Author, book.ISBN)
}

func add_contact(pbook map[string]int, name string, number int) {
	pbook[name] = number
}

func remove_contact(pbook map[string]int, name string) {
	delete(pbook, name)
}

func search_contact(pbook map[string]int, name string) int {
	return pbook[name]
}

func phone_book() {
	phone_book := map[string]int{}
	// add contact
	fmt.Println("add contact ")
	add_contact(phone_book, "John Doe", 1234567890)
	add_contact(phone_book, "Jane Doe", 9876543210)
	for k, v := range phone_book {
		fmt.Println(k, v)
	}
	// remove contact
	remove_contact(phone_book, "John Doe")
	fmt.Println("remove contact ")
	for k, v := range phone_book {
		fmt.Println(k, v)
	}
	// search contact

	num := search_contact(phone_book, "Jane Doe")
	fmt.Println("search contact ", num)
}

func frequency_of_digit(n int) map[int]int {
	freq := make(map[int]int, 0)
	for n != 0 {
		t := n % 10
		n = n / 10
		freq[t]++
	}
	return freq
}
func itersaction_array(arr1 []int, arr2 []int) {
	fmt.Println("intersaction of array")
	mp1 := make(map[int]int)
	mp2 := make(map[int]int)

	len1 := len(arr1)
	len2 := len(arr2)
	mx := max(len1, len2)

	for i := 0; i < mx; i++ {
		mp1[arr1[i]]++
		mp2[arr2[i]]++
	}
	for k, v := range mp1 {
		for k2, v2 := range mp2 {
			if k == k2 {
				if v == v2 {
					fmt.Println(k)
				}
			}
		}
	}
}

func dup_ele(n []int) []int {
	dp_ele := make([]int, 0)
	mp := make(map[int]int, 0)
	for i := 0; i < len(n); i++ {
		mp[n[i]]++
	}
	for k, v := range mp {
		if v > 1 {
			dp_ele = append(dp_ele, k)
		}
	}
	return dp_ele
}

func map_swapper(mp1 map[string]string) map[string]string {
	mp := make(map[string]string, 0)
	for k, v := range mp1 {
		mp[v] = k
	}
	return mp
}

func sec_largest_num(sl []int) int {
	var max1 int = sl[0]
	var max2 int
	for i := 1; i < len(sl); i++ {
		if max1 < sl[i] {
			max2 = max1
			max1 = sl[i]
		}
	}
	return max2
}

func find_mission_num(sl []int) []int {
	slc := make([]int, 0)
	for i := 0; i < len(sl); i++ {
		if i+1 != sl[i] {
			for j := i + 1; j < sl[i]; j++ {
				slc = append(slc, j)
			}
			i = sl[i]
		}
	}
	return slc
}

func split_into_chanks(sl []int, k int) [][]int {
	var chunk [][]int
	for i := 0; i < len(sl); i++ {
		chunk = append(chunk, sl[i:min(i+k, len(sl))])
	}
	return chunk
}

// sliding windows maximun
// func sliding_window_maximun(sl []int, k int)[]int{
// 	// var maxValue []int
// 	for i := 0; i < len(sl)-k; i++ {

// 	}
// }

func typeConvertion(n int, f float32, s string) {
	fmt.Println("int to string : ")
	to_float := float32(n)
	fmt.Printf("value %.2f datatype %T\n", to_float, to_float)
	to_int := int(f)
	fmt.Printf("value %d datatype %T\n", to_int, to_float)
	str_to_float, err := strconv.ParseFloat(s, 64)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Printf("value %.2f datatype %T\n", str_to_float, str_to_float)

}

func swap_two_num(n int, n2 int) {
	fmt.Println("a : ", n)
	fmt.Println("b : ", n2)
	n = n + n2
	n2 = n - n2
	n = n - n2
	fmt.Println("after swap")
	fmt.Println("a : ", n)
	fmt.Println("b : ", n2)
}

func sum_od_digit(n int) int {
	sum := 0
	for n != 0 {
		t := n % 10
		n = n / 10
		sum += t
	}
	return sum
}

// func factorial(n int)int{
// 	fact := 0
// 	for i := 0; i < n; i++ {
// 		fact *= i
// 	}
// 	return fact
// }

func trailing_zeros(n int) int {
	if n < 0 {
		return -1
	}
	var count int = 0
	for i := 5; n/i >= 1; i *= 5 {
		count += int(math.Floor(float64(n) / float64(i)))
	}
	return count
}

func power_of_two_checker(n int) {
	fl := (n&(n-1) == 0)
	fmt.Println(fl)
}

func removeValue(slice []int, val int) []int {
	i := 0
	for _, v := range slice {
		if v != val {
			slice[i] = v
			i++
		}
	}
	slice = slice[:i]
	return slice
}

func main() {
	fmt.Println()
	fmt.Println()
	employee_records()
	fmt.Println()
	fmt.Println()
	student_grade_management()
	fmt.Println()
	fmt.Println()
	area_of_rectangle()
	fmt.Println()
	fmt.Println()
	circle_perimeter_and_area()
	fmt.Println()
	fmt.Println()
	library_book_management()
	fmt.Println()
	fmt.Println()
	phone_book()
	fmt.Println()
	fmt.Println()
	i := 1112223355
	fr := frequency_of_digit(i)
	for k, v := range fr {
		fmt.Println(k, v)
	}
	var arra1 []int = []int{1, 2, 3, 4, 5}
	var arra2 []int = []int{1, 2, 3, 6, 7}
	itersaction_array(arra1[:], arra2[:])
	var sl []int = []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}

	dup := dup_ele(sl)
	fmt.Println("duplicate element :")
	for i := 0; i < len(dup); i++ {
		fmt.Println(dup[i])
	}

	mp3 := map[string]string{
		"nikhil":  "kumar",
		"nikhil2": "kumar2",
		"nikhil3": "kumar3",
	}

	map_swap := map_swapper(mp3)
	for k, v := range map_swap {
		fmt.Println(k, v)
	}
	sec_lar := sec_largest_num(sl)
	fmt.Println(sec_lar)

	slc := []int{1, 2, 5, 6, 7, 9}
	fmt.Println("missing numebr : ", find_mission_num(slc))

	chunks := split_into_chanks(slc, 2)
	fmt.Println("split into chunks ", chunks)
	typeConvertion(132, 132.12, "123456")
	swap_two_num(2, 3)
	sum_od_digit(123456)
	tra := trailing_zeros(100)
	fmt.Println("trailing zeros", tra)
	power_of_two_checker(16)

	fmt.Println("removed value", removeValue(slc, 6))
}
