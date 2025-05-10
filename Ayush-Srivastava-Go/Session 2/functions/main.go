package main

import (
	"fmt"
	"math"
)

//Q.1
func swap(num1 int, num2 int) (int, int) {
	var temp int

	temp = num1
	num1 = num2
	num2 = temp
	return num1, num2
}

// Q.2
func area(radius int) float64 {
	 
	res := math.Pi * float64(radius) * float64(radius)

	return res
}

// Q.3

func celsiustoFar(temp int) int {

	far := (temp * (9/5)) + 32

	return far
}

// Q.4

func concat(name string, surname string) string{
	
	res := name + surname

	return res
}

// Q.5
// TODO:

func dataType()

// Q.6

func sum(num1 int, num2 int)int{
	res := num1+num2

	return res
}

// Q.7
func sumofN(n ...int) int {

	sum := 0

	for i:=0; i<len(n); i++{
		sum += n[i]
	}

	return sum
}

// Q.8
func maxVal(n ...int) int {

	curr := n[0]

	for i:=0; i<len(n); i++{
		if(curr < n[i]){
			curr = n[i]
		}
	}

	return curr
}

// Q.9

func isPalindrome(word string) bool {

	n := len(word)

	for i:=0; i<n/2; i++{
		if word[i] != word[n-1-i]{
			return false
		}
	}
	return true
}

// Q.10

func factorial(num int) int {

	if(num == 0 || num == 1){
		return 1
	}

	return factorial(num-1) * num 

}

// Q.11

func sumofFirstN(num int) int {

	res := 0

	for i:=0; i<num; i++{
		res += (i + 1)
	}

	return res
}

// Q.12
func tableGen(num int){

	for i:=0; i<10; i++{
		fmt.Println((i+1)*num)
	}
}

// Q.13
func fib(n int){

	if(n<=0){
		return
	}

	a,b := 0,1

	fmt.Println(a)

	for i:=1; i<n; i++{
		fmt.Println(b)
		a, b = b, a+b
	}

	fmt.Println()
}

// Q.14
func reverse(s string) string {

	rev := ""

	for i:=0; i<len(s); i++{
		// TODO:
	}

	return rev

}

// Q.15
func isPrime(n int) bool{

	if(n<2){
		return false
	}
	
	for i:=2; i<int(math.Sqrt(float64(n))); i++{
		if(n%i == 0){
			return false
		}
	}

	return true
}

// Q.16
func evenOdd(num int) string {

	if(num%2 == 0){
		return "even"
	}else{
		return "odd"
	}
}

// Q.17
func greatestOfThree(num1 int, num2 int, num3 int) int {

	var temp int

	if(num1 > num2){
		temp = num1
	}else{
		temp = num2
	}

	if(temp > num3){
		return num3
	}else{
		return temp
	}

}

//Q.18
func isLeapYear(year int) bool{
	if(year % 4 == 0){
		return true
	}else{
		return false
	}
}

// Q.19
func gradeCalc(score int) string {
	
	if(score >= 90 && score <= 100){
		return "A"
	}else if(score >= 80 && score < 90){
		return "B"
	}else if(score >= 70 && score < 80){
		return "C"
	}else if(score >= 60 && score < 70){
		return "D"
	}else{
		return "F"
	}
}

// Q.20
func typeOfAlpha(){
	// TODO:
}

