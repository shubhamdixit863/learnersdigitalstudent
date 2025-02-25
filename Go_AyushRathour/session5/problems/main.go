package main

import ("fmt"
"strings")

func main(){
	// phonedirectory()
	// fmt.Println(freq(1112345666))
	// slice1:= []int{1,3,3,4,0,5,3,6}
	// slice2:= []int{4,5,6,7,8,9}
	// fmt.Println(intersection(slice1,slice2))

	// fmt.Println(duplicate([]int{1,2,3,4,5,6,6,6}))
// 	data:= map[string]string{"Ayush":"NIT", "Aman":"IIT", "Aadil":"IIIT"}
// 	fmt.Println(swap(data))

//  fmt.Println(second_largest([]int{1, 2, 5, 6, 76, 89}))
//  fmt.Println(findMissingNumbers([]int{1, 2, 5, 6, 8}, 8))
// fmt.Println(chunkSlice([]int{1,2,3,4,4,6,7,8,9,0}, 3))
// fmt.Println(max([]int{1,2,3,4,4,6,7,8,9,0}, 3))

    // rect := Rectangle{Length: 10.5, Width: 5.2}
	// fmt.Printf("Rectangle with Length %.2f and Width %.2f has an Area of %.2f\n", rect.Length, rect.Width, rect.Area())

	// books := []Book{
	// 	{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", ISBN: "978-0134190440"},
	// 	{Title: "Clean Code", Author: "Robert C. Martin", ISBN: "978-0132350884"},
	// 	{Title: "Design Patterns", Author: "Erich Gamma", ISBN: "978-0201633610"},
	// }


	// searchTitle := "clean code"
	// foundBook := FindBook(books, searchTitle)

	
	// if foundBook != nil {
	// 	fmt.Printf("Book Found: \nTitle: %s\nAuthor: %s\nISBN: %s\n",
	// 		foundBook.Title, foundBook.Author, foundBook.ISBN)
	// } else {
	// 	fmt.Println("Book not found.")
	// }

	fmt.Println(remove_value([]int{1, 3, 4, 5, 4, 6, 7}, 4))



}


//1. Create a struct for an Employee with fields like ID, Name, and Salary. Write a function to give a salary hike to employees earning below a certain threshold.

type Employee struct{
	id int
	name string
	Salary float64

}

func giveSalaryHike(employees []Employee, threshold float64, hikePercentage float64) {
	for i := range employees {
		if employees[i].Salary < threshold {
			employees[i].Salary += employees[i].Salary * (hikePercentage / 100)
		}
	}
}


//2. Define a Student struct with Name, Marks, and Grade. Write a function to assign grades based on marks.

type Student struct{
	Name string
	Marks int
	Grade string
}

func assignGrade(students []Student) {
	for i := range students {
		switch {
		case students[i].Marks >= 90:
			students[i].Grade = "A+"
		case students[i].Marks >= 80:
			students[i].Grade = "A"
		case students[i].Marks >= 70:
			students[i].Grade = "B"
		case students[i].Marks >= 60:
			students[i].Grade = "C"
		case students[i].Marks >= 50:
			students[i].Grade = "D"
		default:
			students[i].Grade = "F"
		}
	}
}

//3.Create a Rectangle struct with Length and Width. Write a method to calculate the area.
type Rectangle struct{
	Length float32
	Width float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

//4. Define a Circle struct with a Radius field. Write methods to calculate the circumference and area.
type Circle struct{
	Radius float64

}
func (r Circle) CircleArea() float64{
	return 22.7*r.Radius*r.Radius

}

//5. Create a Book struct with fields Title, Author, and ISBN. Write a function to search for a book by title in a slice of books.
type Book struct{
	Title string
	Author string
	ISBN string
}

func FindBook(books[]Book , title string) *Book{

	for _, book := range books {
		if strings.EqualFold(book.Title, title) {
			return &book
		}
	}
	return nil

}



//6.
func phonedirectory(){
	phonebook := make(map[string]int64)
	for{
	fmt.Println("1. ADD contact, 2. DELETE contact, 3. SEARCH 4.Exit")
	var n int
	fmt.Scanln(&n)
	switch n {
	case 1:
		AddContact(phonebook)
	case 2:
		DeleteContact(phonebook)
	case 3:
		Search(phonebook)
	case 4:
		return
	default:
		fmt.Println(" Enter a valid choice ")
		return
	}
}

}

func AddContact(phonebook map[string]int64) {
	fmt.Printf("Enter name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Enter number: ")
	var number int64
	fmt.Scanln(&number)

	if phonebook[name] == 0 {
		phonebook[name] = number
	}else{
		fmt.Println("Name already exists")
	}
}

func DeleteContact(phonebook map[string]int64){
	fmt.Printf("Enter name : ")
	var name string
	fmt.Println(&name)
	delete(phonebook, name)
}

func Search(phonebook map[string]int64){
	fmt.Printf("Enter name : ")
	var name string
	fmt.Println(&name)
	for n, no:= range phonebook{
		if n==name{
			fmt.Printf("Phone No. is : %d", no)
		}
	}


}


//7. Given an integer, return a map containing the frequency of each digit.
func freq(num int) map[int]int{
	num_map:=make(map[int]int)
	for num>0{
		d:=num%10
		if num_map[d]==0{
			num_map[d]=1
		}else{
			num_map[d]++
		}
		num=num/10
	}
   return num_map
}


//8. Write a function that finds the intersection of two integer slices using maps.
func intersection(slice1, slice2 []int)[]int{
	m := make(map[int]bool)
	var result []int

	for _, num := range slice1 {
		m[num] = true
	}
	for _, num := range slice2 {
		if m[num] {
			result = append(result, num)
			delete(m, num) 
		}
	}

	return result
}


//9. Given a slice of integers, use a map to find and return the list of duplicates.
func duplicate(slice[]int)[]int{
	m:=make(map[int]int)
	var result []int
	 
	for i:=0; i<len(slice); i++{
		if m[slice[i]]==1{
			result= append(result, slice[i])
			m[slice[i]]++
		}else{
			m[slice[i]]=1
		}
	}
	return result
}


//10. Given a map of string keys and string values, create a new map where keys and values are swapped.

func swap(data map[string]string) map[string]string{
	swaped_data:=make(map[string]string)
	for key,value:=range data{
		swaped_data[value]=key
	}
	return swaped_data
}


//11. Write a function to find the second largest element in a slice of integers.

func second_largest(num []int)int{
	large:=num[0]
	second_large:=num[0]

	for i:=1; i<len(num); i++{

		if num[i]>large{
			second_large=large
			large=num[i]
		}else{
			if num[i]>second_large{
				second_large=num[i]
			}
		}
	}
return second_large
}

//12. Given a slice of integers from 1 to n with some numbers missing, find all missing numbers.
func findMissingNumbers(nums []int, n int) []int {
	exists := make(map[int]bool)
	var missing []int
	for _, num := range nums {
		exists[num] = true
	}
	for i := 1; i <= n; i++ {
		if !exists[i] {
			missing = append(missing, i)
		}
	}

	return missing
}

//13. Write a function to split a slice into smaller chunks of size k.

func chunkSlice(slice []int, k int) [][]int {
	if k <= 0 {
		return nil 
	}

	var chunks [][]int
	for i := 0; i < len(slice); i += k {
		end := i + k
		if end > len(slice) {
			end = len(slice) 
			
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}


//14. Implement a function to find the maximum in every contiguous sub-slice of length k.

func max(num []int, k int) []int {
	var results[]int
	for i := 0; i <=len(num)-k; i ++ {
		end := i + k
		if end > len(num) {
			end = len(num) 
			
		}
		results = append(results, maximum(num[i:end]))
	}
	return results
}
func maximum(data []int )int{
	max:=data[0]
	for i:=0; i<len(data); i++{
       if data[i]>max{
		max=data[i]
	   }
	}
	return max
}

//15. Given a slice and a value, remove all occurrences of that value without creating a new slice.

func remove_value(num[]int , k int)[]int{
    a:=0
	for i:=0; i<len(num); i++{
      if num[i]!=k{
		num[a]=num[i]
        a++
	  }
	}
	num =num[:a]

return num
}

//16.Write functions to convert between basic data types (int to string, string to float, etc.).

//17. Create a function that calculates the sum of all digits in an integer.
func sum(num int) int{
	sum:=0
	for num>0{
		sum=sum +num%10
		num=num/10
	}
  return sum
}

//18. Implement a function to swap two integers without using a third variable.
func swap_num(){
		a, b := 5, 10
		fmt.Println("Before swapping: ", a, b)
		a=a+b
		b=a-b
		a=a-b
		fmt.Println("After swapping: ", a, b)
	   
}

//19. Write a function to find the number of trailing zeroes in the factorial of a number.

func TrailingZeroes(n int) int {
	count := 0
	for i := 5; n/i > 0; i *= 5 {
		count += n / i
	}
	return count
}

//20.  Write a function to check if a given integer is a power of two.

func isPowerOf2(n int)bool{
	return n > 0 && (n & (n - 1)) == 0

}