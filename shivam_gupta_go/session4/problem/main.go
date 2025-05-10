package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a []int
	a = append(a, 5)
	a = append(a, 6)
	a = append(a, 8)
	a = append(a, 10)
	a = append(a, 10)

	// fmt.Println("reverse array")

	// reverseArray(a)

	// sum(a)
	// maaximun(a)
	// Rotate(a, 3)
	// RemoveDublicate(a)
	// Even(a)
	// Occurrences(a)
	// Frequency()
	// CharacterFrequency()
	// TwoSum(a)
	// Duplicates(a)
	// firstUniq()
	// fibbo()
	// palindrome()
	Missing()
}

func reverseArray(a []int) {
	var b []int
	for i := len(a) - 1; i > 0; i-- {
		b = append(b, a[i])
	}
	fmt.Println(b)
}

// Sum of Elements in a Slice:

// Create a function that takes a slice of integers and returns the sum of all elements.

func sum(a []int) {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}

// Find Maximum and Minimum in an Array:

// Given an integer array, find and return the maximum and minimum elements.

func maaximun(a []int) {
	maxi := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > maxi {
			maxi = a[i]
		}
	}
	fmt.Println(maxi)
}

// Remove Duplicates from a Slice:

// Write a function to remove all duplicate integers from a slice.

func RemoveDublicate(a []int) {
	var b []int
	for i := 0; i < len(a); i++ {

		for j := i + 1; j < len(a); j++ {
			if a[j] == a[i] {
				i++
				continue
			}
		}
		b = append(b, a[i])
	}
	fmt.Println(b)
}

// Rotate Array by k Steps:

// Rotate the elements of an array to the right by k steps.

func Rotate(a []int, k int) {
	b := k % len(a)
	var c []int
	var d []int
	for i := len(a) - b; i < len(a); i++ {
		c = append(c, a[i])
	}
	for i := 0; i < len(a)-b; i++ {
		d = append(d, a[i])
	}

	e := append(c, d...)
	fmt.Println(e)
}

func Even(a []int) {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			count++
		}
	}
	fmt.Println(count)
}

// Count Occurrences of an Element:

// Write a function that counts how many times a specific integer appears in a slice.

func Occurrences(a []int) {
	mp := make(map[int]int)
	for _, num := range a {
		mp[num]++
	}
	fmt.Println(mp)
}


// Word Frequency Counter:

// Given a string, count the frequency of each word using a map.

func Frequency(){
	fmt.Println("enter sentences")
	reader := bufio.NewReader(os.Stdin)
	input,_ :=reader.ReadString('\n')
	words:=strings.Fields(input) 

	mp:=make(map[string]int)

	for _,num:= range words{
		mp[num]++
	}
	fmt.Println(mp)
}


// Character Frequency in a String:

// Write a function that returns a map with the count of each character in a given string.

func CharacterFrequency(){
	fmt.Println("enter one word")
  var s string
	fmt.Scanln(&s)
  mp:=make(map[rune]int)
	for _,i:=range s{
   mp[i]++
	}
	for key,value:=range mp{
		fmt.Printf("%c : %d\n",key,value)
	}
}


// Two Sum Problem:

// Given an array and a target integer, find the indices of the two numbers that add up to the target using a map.

func TwoSum(a[]int){
	fmt.Println("two sum problem")
	var inp  int 
	fmt.Println("enter target element")
	fmt.Scanln(&inp)

	mp := make(map[int]int )

	for i,num:=range(a){
     compliment:=inp-num
		 if index,value :=mp[compliment]; value {
      fmt.Printf("%d %d and indexes are %d , %d",num,compliment,i,index)
			return
		 }
    
		mp[num]=i
        
		 
	}

  fmt.Println("not found")
}



// Check for Duplicates in an Array:

// Determine if an array contains any duplicate elements using a map.


func Duplicates(a[]int){
	fmt.Println("check dublicates ")
	
	mp := make(map[int]int)

	for _,num:= range (a){
     mp[num]++
	}
   
	for key,val:=range(mp){
		if(val>1){
		fmt.Println(key)}
	}
	
}

// First Unique Character in a String:

// Find the index of the first non-repeating character in a string using a map.



func firstUniq(){
	fmt.Println("first unique character ")
	fmt.Println("enter string")
	var s string 
	fmt.Scanln(&s)
	mp:=make(map[rune]int)
	for _,num:=range(s){
		mp[num]++
	}
	for _,val:=range(s){
		if mp[val]==1 {
			fmt.Printf("%c",val)
			return
		}
	}
}


// Fibonacci Sequence Generator:

// Write a function that generates the first n Fibonacci numbers using a for loop.

func fibbo(){
	fmt.Println("enter number to generate")
	var a int
	fmt.Scanln(&a)
	b:=0
	c:=1
	for i:=0;i<a;i++{
		fmt.Println(b)
		next:=b+c
		b=c
		c=next
		
	} 
}


// Palindrome Number Checker:

// Write a function to check if a given integer is a palindrome (without converting it to a string).

func palindrome(){
	fmt.Println("palindrome")
	fmt.Println("enter string")
	var s string
	fmt.Scanln(&s)
	i:=0
	e:=len(s)-1
	for(i<e) {
     if(s[i]!=s[e]){
			fmt.Println("not palindrome")
			return
		 }
			i++
			e--
		 

	}
	fmt.Println("palindrome")
}

// Find Missing Number in a Range:

// Given an array containing n-1 distinct integers from 1 to n, find the missing number.


func Missing() {
	fmt.Println("missing number ")
	fmt.Println("enter size of array")
	var n int 
	fmt.Scanln(&n)
	arr:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Println("enter ",i+1,"element")
		var a int 
		fmt.Scanln(&a)
		arr = append(arr, a)
	}
  sum:=0
	for i:=0;i<len(arr);i++{
		sum+=arr[i]
	}
	totalSum := (n*(n+1))/2
	missingNum:=sum-totalSum
	fmt.Println("missing number : ", missingNum)
}
