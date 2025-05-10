package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

}

// Q.1
type Employee struct {
	id     int
	name   string
	salary int
}

func giveHike(employees []Employee, value int, hike int) {

	for _, data := range employees {
		if data.salary < value {
			data.salary += hike
		}
	}
}

// Q.2
type Student struct {
	name  string
	marks int
	grade string
}

func assignGrades(students []Student) {

	for _, data := range students {
		marks := data.marks

		if marks > 90 && marks <= 100 {
			data.grade = "A"
		} else if marks > 80 && marks <= 90 {
			data.grade = "B"
		} else if marks > 70 && marks <= 80 {
			data.grade = "C"
		} else {
			data.grade = "D"
		}
	}
}

// Q.3
type Rectengle struct {
	width  float64
	length float64
}

func AreaRectengle(r Rectengle) float64 {
	return (r.width) * (r.length)
}

// Q.4
type Circle struct {
	radius float64
}

func AreaCircle(c Circle) float64 {
	return math.Pi * c.radius * c.radius
}

// Q.5
type Book struct{
	Title string
	Author string
	ISBN string
}

func searchBook(books []Book, title string) string {

	for _, book := range books{

		if(book.Title == title){
			return book.ISBN
		}
	}

	return "not found!"
}

// Q.6

var directory = make(map[string]string)

func addContact(name string, number string){

	_, ok := directory[name]

	if(!ok){
		directory[name] = number
	}
}

func deleteContact(name string){

	_, ok := directory[name]

	if(ok){
		delete(directory, name)
	}
}

func searchContact(name string)(string){

	number,ok := directory[name]
	
	if(ok){
		return number
	}

	return "Not found!"
}

// Q.7
func getFreq(num int) map[rune]int {

	mp := make(map[rune]int)

	s := strconv.Itoa(num)

	for _, ch := range s{
		mp[ch]++
	}

	return mp
}

// Q.8
func intersection(a1 []int, a2 []int) []int {

	mp := make(map[int]int)

	res := make([]int, 0)

	for _, val := range a1{
		mp[val]++
	}

	for _, val := range a2{
		
		data, ok := mp[val]

		if(ok){
			if(data != 0){
				res = append(res, data)
				mp[val]--
			}
		}
	}
	return res
}

// Q.9
func findDuplicates(arr []int) []int {

	res := make([]int, 0)

	mp := make(map[int]int)

	for _, data := range arr{
		mp[data]++
	}

	for k, v := range mp{
		if (v >= 2){
			res = append(res, k)
		}
	}

	return res
}

// Q.10
func swapMap(mp map[string]string) map[string]string{

	newMap := make(map[string]string)

	for k, v := range mp{
		newMap[v] = k
	}

	return newMap
}

// Q.11
func secondLargest(arr []int) int {

	n := len(arr)

	largest := -1
	secLargest := -1

	for i:=0; i<n; i++{
		if(arr[i] > largest){
			largest = arr[i]
		}
	}

	for i:=0; i<n; i++{
		if(arr[i] > secLargest && arr[i] != largest){
			secLargest = arr[i]
		}
	}

	return secLargest
}

// Q.12
func findMissings(arr []int, n int) []int {

	missing := make([]int, 0)

	mp := make(map[int]bool)

	for _, val := range arr{
		mp[val] = true
	}

	for i:=1; i<=n; i++{
		
		exists, _ := mp[i]

		if(!exists){
			missing = append(missing, i)
		}
	}

	return missing
}

// Q.13
func splitIntoK(str string, k int) []string{

	n := len(str)

	res := make([]string, 0)

	chunks := n/k

	for i:=0; i<k; i++{

		temp := ""

		for j:=0; j<chunks; j++{
			temp += string(str[i*chunks+j])
		}

		res = append(res, temp)
	}

	return res
}

// Q.14
// func getMaxVal(arr []int) []int{

// }


// Q.15
func removeVal(arr []int, k int) []int{

	n := len(arr)

	idx := 0

	for i:=0; i<n; i++{
		if(arr[i] != k){
			arr[idx] = arr[i] 
		}
		idx++
	}

	return arr[0:idx]
}

// Q.16
func conversion(data string, from string, to string)(string, error){

	switch from{
	case "int":
		i, err := strconv.Atoi(data)
		if(err != nil){
			return "", err
		}

		switch to{
		case "string":
			return strconv.Itoa(i), nil
		case "float":
			return fmt.Sprintf("%2f", float64(i)), nil
		case "bool":
			return strconv.FormatBool(i != 0), nil
		}
	}

	//TODO:

	return "Similarly", nil
}

// Q.17
func sumDigits(num int) int {

	sum := 0

	for (num > 0){
		sum += num%10
		num /= 10
	}

	return sum
}

// Q.18
func swapNums(num1 int, num2 int)(int, int){

	num1, num2 = num2, num1

	return num1, num2
}

// Q.19
func trailingZeros(n int)int{

	cnt := 0

	for n >= 5{
		n /= 5
		cnt += n
	}

	return cnt
}

// Q.20
func isPowof2(n int) bool {

	if(n <= 0){
		return false
	}

	return (n & (n-1) == 0)
}