package main

import (
	"fmt"
	"math"
	"strconv"
)

// 1.	Employee Records:

// Create a struct for an Employee with fields like ID, Name, and Salary. Write a function to give a salary hike to employees earning below a certain threshold.

type Employee struct {
	id     int
	name   string
	salary int
}

func main() {
	//  salaryhike()
	// studentGrade()
	// RectangleArea()
	// CirclePerimeterAndArea()
	// phoneDirectory()
	// bookManagement()
	// frequency()
	// intersection()
	// dublicates()
	// Mapswapper()
	// secondLargest()
	// MissingNumber()
	// Slicing()
	// typeConversion()
	// calSum()
	// swap()
	// trailing()
	// PowerChecker()
	removeOcc()
}

func salaryhike() {
	var employee1 Employee
	employee1.id = 1
	employee1.name = "shivam"
	employee1.salary = 2300
	var employee2 Employee
	employee2.id = 2
	employee2.name = "nikhil"
	employee2.salary = 200

	var employee3 Employee
	employee3.id = 3
	employee3.name = "ayush"
	employee3.salary = 20
	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)

	fmt.Println("after hike salary : ")
	employee3.salary += 5000
	fmt.Println(employee3)
}

// Student Grade Management:

// Define a Student struct with Name, Marks, and Grade. Write a function to assign grades based on marks.

func studentGrade() {
	type student struct {
		Name  string
		Marks int
		Grade string
	}

	var student1 student
	student1.Name = "shivam"

	student1.Marks = 95

	var student2 student
	student2.Name = "nikhil"

	student2.Marks = 60

	var student3 student
	student3.Name = "ayush"
	student3.Marks = 75

	var student4 student
	student4.Name = "sanskriti"
	student4.Marks = 55

	var slice []student
	slice = append(slice, student1)
	slice = append(slice, student2)
	slice = append(slice, student3)
	slice = append(slice, student4)

	for i := range slice {
		if slice[i].Marks > 90 {
			slice[i].Grade = "A"
		} else if slice[i].Marks > 80 {
			slice[i].Grade = "B"
		} else if slice[i].Marks > 70 {
			slice[i].Grade = "C"
		} else if slice[i].Marks > 40 {
			slice[i].Grade = "D"
		}

	}
	fmt.Println(slice)
}

// Rectangle Area Calculator:

// Create a Rectangle struct with Length and Width. Write a method to calculate the area

func RectangleArea() {
	type Rectangle struct {
		Length int
		width  int
	}

	var Rectangle1 Rectangle
	Rectangle1.Length = 4
	Rectangle1.width = 5

	fmt.Println(Rectangle1)
	area := Rectangle1.width * Rectangle1.Length
	fmt.Println("area of rectangle : ", area)
}


// Circle Perimeter and Area:

// Define a Circle struct with a Radius field. Write methods to calculate the circumference and area.

func CirclePerimeterAndArea(){
	type circle struct{
		Radius int 
	}
	var circle1 circle
	circle1.Radius=5
	circumference:=2*math.Pi*float64(circle1.Radius)
	fmt.Println("circumference : ",circumference)
	area:=math.Pi*float64(circle1.Radius)*float64(circle1.Radius)
	fmt.Println("area :",area)
}


// Library Book Management:

// Create a Book struct with fields Title, Author, and ISBN. Write a function to search for a book by title in a slice of books.


func bookManagement(){
	type book struct {
		Title string
		Author string
		ISBN string
	}
	var book1 book
	book1.Title="go1"
	book1.Author="shivam1"
	book1.ISBN="1234567"


	var book2 book
	book2.Title="go2"
	book2.Author="shivam2"
	book2.ISBN="12345678"

	var book3 book
	book3.Title="go3"
	book3.Author="shivam3"
	book3.ISBN="123456789"
  
 var slice []book
 slice = append(slice, book1,book2,book3) 

	fmt.Println("enter title of book ")
	var s string
	fmt.Scanln(&s)

	for i :=range (slice){
		if(s==slice[i].Title){
			fmt.Println("found! ",slice[i].Title)
		}else{
			fmt.Println("not found")
		}
	}

}



// Phone Directory:

// Create a phone directory using a map where names are keys and phone numbers are values. Implement functions to add, delete, and search for contacts

func phoneDirectory (){
	mp:=make(map[string]int)
	
	mp["nikhil"]=1234567890
	mp["ayush"]= 9877654321

	fmt.Println("enter name and phone numbe to add (i.e. name phone number)")
	var s string
	var n int
  fmt.Scanln(&s,&n)
	mp[s]=n
  fmt.Println(mp)
	fmt.Println("enter name to delete")
	var name string
	fmt.Scanln(&name)
	for key,_:=range mp{
		if(key==name){
			fmt.Println("deleting.....",key)
			delete(mp,key)
		}else {
			fmt.Println("error!")
		}
	}
  fmt.Println("enter name to search")
		var name1 string
		fmt.Scanln(&name1)
		for key,_:=range(mp){
			if(name1==key){
				fmt.Println("found! : ",key)
			}else {
				fmt.Println("not found")
			}
		}
	}


	// Frequency of Digits in an Integer:

	// Given an integer, return a map containing the frequency of each digit.

	func frequency(){
		mp:=make(map[int]int)
		fmt.Println(" digit frequency checker ")
		fmt.Println("enter integer")
		var n int 
	  fmt.Scanln(&n)
		for n!=0{
      t:=n%10
			n=n/10
			mp[t]++
		}
	 fmt.Println(mp)
	}

	// Intersection of Two Arrays:

	// Write a function that finds the intersection of two integer slices using maps.
	
	
func intersection(){
  fmt.Println("intersection of array")
	arr1:=[]int{1,3,4,5}
	arr2:=[]int{4,5,6,7}

 fmt.Println("array 1 : ",arr1)
 fmt.Println("array 2 : ",arr2)

 mp := make(map[int]int)

 for i:=range arr1{
	mp[arr1[i]]++
 }
 for i:=range arr2{
	mp[arr2[i]]++
 }
 fmt.Println("intersection using map : ")

 for key,val:= range(mp){
     if(val>=2){
			fmt.Println(key)
		 }
 
 }

}


// Find Duplicate Elements:

// Given a slice of integers, use a map to find and return the list of duplicates.

func dublicates (){
	fmt.Println("dublicates integer ")
	slice:=[]int{2,45,12,56,2,12,23,78}
	fmt.Println("slice : ",slice)
	mp:=make(map[int]int)
	for i := range (slice){
		mp[slice[i]]++
	}
	fmt.Println("dublicates integers")
	for key,val := range (mp){
		if(val>=2){
			fmt.Println(key)
		}
	}
}

// Map Value Swapper:

// Given a map of string keys and string values, create a new map where keys and values are swapped.

func Mapswapper(){
	fmt.Println("map swapper")
	mp1:= make(map[string]string)
	mp2 := make(map[string]string)
  mp1["shivam"]="gupta"
	mp1["nikhil"]="kumar"
	for key,val:=range(mp1){
		mp2[val]=key
	}

	fmt.Println("map 1 : ",mp1)
	fmt.Println("map 2 : ",mp2)
}


// Second Largest Element in a Slice:

// Write a function to find the second largest element in a slice of integers.


func secondLargest(){
  fmt.Println("second largest")
	slice:=[]int{1,34,5,12,67,23,78,56}
	var maxi int 
	for i:=range slice {
		if(maxi<slice[i]){
			maxi=slice[i]
		}
	}
	maxi2:=0
	for i:=range slice{
		if(maxi2<slice[i] && maxi>slice[i] ){
			maxi2=slice[i]
		}
	}
	fmt.Println("slice : ",slice)
	fmt.Println("second largest : ",maxi2)
}


// Find Missing Numbers in a Sequence:

// Given a slice of integers from 1 to n with some numbers missing, find all missing numbers.


func MissingNumber(){
	fmt.Println("missing number ")
	slice:=[]int{1,3,4,6,7,8,9}

	var slice2 []int

	Length:=len(slice)
  n := slice[Length-1]

	mp:=make(map[int]bool)

  

	for i:=range slice{
		mp[slice[i]]=true
	}

	for i:=1;i<=n;i++{
		if !mp[i]{
			slice2=append(slice2, i)
		}
	}
	
 fmt.Println("slice : ",slice)
 fmt.Println("missing number : ",slice2)
}


// Split a Slice into Chunks:

// Write a function to split a slice into smaller chunks of size k.


func Slicing(){
	fmt.Println("slicing .....")
	fmt.Println("enter value of k for slicing")
	var k int 
	fmt.Scanln(&k)
  arr:=[]int{1,2,3,4,6,7,8,9,12,34,5}
	
	var slice2 [][]int
	for i:=0;i<len(arr);i=i+k{
    end:=i+k
		if(end>len(arr)){
			end=len(arr)
		}
		slice2 = append(slice2,arr[i:end])
	}
	fmt.Println("sliced array : ",slice2)
}




// Type Conversion Utility:

// Write functions to convert between basic data types (int to string, string to float, etc.).

func typeConversion(){
	a:=2.3
	b:=int(a)
	c:=2
	d:=float32(c)
	e:="123445"
	f,_:=strconv.Atoi(e)
	fmt.Println("type conversion")
	fmt.Printf("a: %T, value: %f → after conversion → b: %T, value: %d\n", a, a, b, b)
	fmt.Printf("c: %T, value: %d → after conversion → d: %T, value: %f\n", c, c, d, d)
	fmt.Printf("e: %T, value: %v → after conversion → f: %T, value: %v\n", e, e, f, f)}



	// Sum of Digits of an Integer:

	// Create a function that calculates the sum of all digits in an integer.

func calSum(){
	fmt.Println("calculates the sum of all digits in an integer")
	fmt.Println("enter integer")
	var n int 
	fmt.Scanln(&n)
	sum:=0
	for n>0{
    t:=n%10
		n=n/10
		sum+=t
	}
	fmt.Println("sum : ",sum)
}

// Swap Two Variables Without a Temp Variable:

// Implement a function to swap two integers without using a third variable.


func swap(){
	fmt.Println("swap")
	fmt.Println("enter 1st number ")
	var n int 
	fmt.Scanln(&n)
	fmt.Println("enter 2nd number ")
	var n2 int 
	fmt.Scanln(&n2)
  n,n2=n2,n
  fmt.Println("swap numbers : ",n,n2)
}


// Trailing Zeroes in Factorial:

// Write a function to find the number of trailing zeroes in the factorial of a number.

func trailing(){
	println("trailing zero")
	fmt.Println("enter value")
	var n int 
	fmt.Scanln(&n)
	count:=0
	for i:=5;i<=5;i*=5{
		count+=n/i
	}
	fmt.Println("trailing zeros : ",count)
}

// Power of Two Checker:

// Write a function to check if a given integer is a power of two.

func PowerChecker(){
	fmt.Println("power checker ")
	fmt.Println("enter integer value")
	var n int 
	fmt.Scanln(&n)
	if(n==0){
		fmt.Println("false")
	}

	flag := (n & (n-1) == 0)

	fmt.Println(flag)
}

// Remove All Occurrences of an Element:

// Given a slice and a value, remove all occurrences of that value without creating a new slice.


func removeOcc(){
	fmt.Println("remove all occurrences of that value without creating a new slice")
	arr:=[]int{1,2,3,4,5,6,12,1}
	fmt.Println("arr : ",arr)
	fmt.Println("enter value to delete")
	var n int 
	fmt.Scanln(&n)
	j:=0
  for i:=0;i<len(arr);i++{
		if(arr[i]!=n){
    arr[j]=arr[i]
		j++
		}
	}
	arr2:=arr[:j]
	fmt.Println(arr2)
}