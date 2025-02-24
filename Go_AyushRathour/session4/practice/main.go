package main

import (
	"fmt"
)

func main() {


	fmt.Printf("Enter the numbers : ")
	num:=input()

	// // reverseArray(num)
	// // fmt.Println("----------------------- ")
	// sum(num)
	// fmt.Println("----------------------- ")
	// MaxMin(num)
	// fmt.Println("----------------------- ")
	// removeDuplicates(num)
	// fmt.Println("----------------------- ")
	rotateArray(num)
	// fmt.Println("----------------------- ")
	// findEven(num)
	// fmt.Println("----------------------- ")
	// countOccurance(num)
	// fmt.Println("----------------------- ")
	// wordFeq()
	// fmt.Println("----------------------- ")
	// charFeq()
	// fmt.Println("----------------------- ")
	// twoSum(num)
	// fmt.Println("----------------------- ")
	// indexOfNonrepeatingChar()
	// fmt.Println("----------------------- ")
	// fib()
	// fmt.Println("----------------------- ")
	// isPalindrome()
	// fmt.Println("----------------------- ")
	// MissingNum()
	// fmt.Println("----------------------- ")
	


}

func input()[]int{
	var num []int
	var n int
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		num = append(num, n)
	}
	return num
}
  


func reverse(num []int, start int , end int) []int {
	n:=end-start
	
	for i := 0; i <n/2; i++ {
		num[i+start], num[end-i-1] = num[end-i-1], num[i+start]
	}

	return num
}


//1.
func reverseArray(num[]int){
	fmt.Println("Reversed Array : ")

	l:=len(num)
	reverse(num,0,l)
	fmt.Println(num)
}

//2.
func sum(num[]int) {

	sum := 0
	for i:=0;i<len(num);i++{
		sum+=num[i]
	}
	fmt.Printf("sum of numbers in array is %d: ", sum)
}

//3.
func MaxMin(num[]int){
    
	max:=num[0]
	min:=num[0]
	for i:=1;i<len(num);i++{
		if num[i]>max{
			max=num[i]
		}
		if num[i]<min{
			min=num[i]
		}
	}
	fmt.Println("Max:",max)
	fmt.Println("Min:",min)
}

//4.

func removeDuplicates(num[]int){
	
	m:=make(map[int]bool)
	for i:=0;i<len(num);i++{
		if !m[num[i]]{
			m[num[i]]=true
		}
	}
	for k:=range m{
		fmt.Println(k)
	}
}

//5.

func rotateArray(num []int){
	
	fmt.Println("Enter the number of rotations:")
	var n int
	fmt.Scanf("%d",&n)
	l:=len(num)
	reverse(num,0,n)
	fmt.Println(num)
	reverse(num,n,l)
	fmt.Println(num)
	reverse(num,0,l)
	fmt.Println(num)
}


//6.
func findEven(num []int){
	fmt.Println("Even numbers in the array:")
	for i:=0;i<len(num);i++{
		if num[i]%2==0{
			fmt.Println(num[i])
		}
	}
}


//7.

func countOccurance(num []int){
	fmt.Println("Occurance of numbers in the array:")
	m:=make(map[int]int)
	for i:=0;i<len(num);i++{
		if m[num[i]]!=0{
			m[num[i]]++
		}else{
			m[num[i]]=1
		}
	}
	for k,v:=range m{
		fmt.Println(k,v)
	}
}

//8.
func wordFeq(){

	fmt.Println("Enter a Sentence: ")
	var sen []string
	var s string
	for {
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			break
		}
		sen = append(sen, s)
	}
	fmt.Printf("\n")

	m:=make(map[string]int)
	for i:=0;i<len(sen);i++{
		if m[sen[i]]!=0{
			m[sen[i]]++
		}else{
			m[sen[i]]=1
		}
	}
	for k,v:=range m{
		fmt.Println(k,v)
	}


}






//9.
func charFeq(){
	
	m:=make(map[string]int)
	var s string
	// str:="hello"
	// str1:="こんにちは"
	// fmt.Println([]byte(str))
	// fmt.Println([]rune(str))
	// fmt.Println([]byte(str1))
	// fmt.Println([]rune(str1))
	fmt.Println("Enter the string: ")
	fmt.Scanf("%s",&s)
	fmt.Println(len(s))
	for i:=0;i<len(s);i++{
		if m[string(s[i])]!=0{
			m[string(s[i])]=m[string(s[i])]+1
		}else{
			m[string(s[i])]=1
		}
	}
	for k,v:=range m{
		fmt.Println(k,v)
	}
}

//10.
func twoSum(num[] int){
	var a int
	fmt.Println("Enter the target sum: ")
	fmt.Scanf("%d",&a)
	m:=make(map[int]int)
	for i:=0;i<len(num);i++{
		if m[a-num[i]]!=0{
			fmt.Println(m[a-num[i]],i)
		}else{
			m[num[i]]=i
		}
	}
}


// 11. same as 7.

// 12.


func indexOfNonrepeatingChar(){
	
	m:=make(map[string]int)
	var s string
	fmt.Println("Enter the string:")
	fmt.Scanf("%s",&s)
	for i:=0;i<len(s);i++{
		if m[string(s[i])]!=0{
			m[string(s[i])]=m[string(s[i])]+1
		}else{
			m[string(s[i])]=1
		}
	}
	for i:=0;i<len(s);i++{
		if m[string(s[i])]==1{
			fmt.Println(i)
			break
		}
	}
}

// 13. 


func fib(){
	var n int
	fmt.Println("Enter the number of fibonacci numbers:")
	fmt.Scanf("%d",&n)
    a, b:=0, 1
	for i:=0;i<n;i++{
		fmt.Println(a)
		a,b=b,a+b
	}

}


//14.


func isPalindrome(){
	var n int
	fmt.Println("Enter the number:")
	fmt.Scanf("%d",&n)
	rev:=0
	temp:=n
	for n>0{
		rev=rev*10+n%10
		n=n/10
	}
	if rev==temp{
		fmt.Println("Palindrome")
	}else{
		fmt.Println("Not Palindrome")
	}
}

// 15.
// Given an array containing n-1 distinct integers from 1 to n, find the missing number./
func MissingNum(){
	var n int
	fmt.Println("Enter the number of elements:")
	fmt.Scanf("%d",&n)
	var num []int
	for i:=0;i<n;i++{
		var a int
		fmt.Scanf("%d",&a)
		num=append(num,a)
	}
	sum:=0
	for i:=0;i<n;i++{
		sum+=num[i]
	}
	fmt.Println(n*(n+1)/2-sum)
}


