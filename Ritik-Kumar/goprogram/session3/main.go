package main

import (
	"fmt"
	"reflect"
	"strconv"
	"math"
)
func main(){
	var st1 string="101"
	var st2 string="506"
	c1,err:=strconv.Atoi(st1)
	if(err!=nil){
		fmt.Println("error",err)
	}
	fmt.Println(reflect.TypeOf(c1))
	c2,err:=strconv.Atoi(st2)
	if(err!=nil){
		fmt.Println("error",err)
	}
	fmt.Println(reflect.TypeOf(c2))
	sum:=c1+c2
	fmt.Println(sum)

	var i1 int =5
	f:=float64(i1)
	fmt.Println(reflect.TypeOf(f))

	var b bool= true
	var e1 int
	if(b){
		e1=1
	}else{
		e1=0
	}
	println(e1)
	
	var i2 float32= 8.2
	i3:=int32(i2)
	if(i3%2==0){
		fmt.Println("even")
	}else{
		fmt.Println("odd")
	}


	var cha4 string
	fmt.Println("enter hexadecimal")
	fmt.Scanln(&cha4)
	var sum4 int
	var ans int=0
	for i:=0;i<len(cha4);i++{
		inx,arr:=strconv.Atoi(string(cha4[i]))
		if(arr!=nil){
			if(cha4[i]=='A'){
				sum4=int(math.Pow(16,float64(len(cha4)-1-i)))*10
			}else if(cha4[i]=='B'){
				sum4=int(math.Pow(16,float64(len(cha4)-1-i)))*11
			}else if(cha4[i]=='C'){
				sum4=int(math.Pow(16,float64(len(cha4)-1-i)))*12
			}else if(cha4[i]=='D'){
				sum4=int(math.Pow(16,float64(len(cha4)-1-i)))*13
			}else if(cha4[i]=='E'){
				sum4=int(math.Pow(16,float64(len(cha4)-1-i)))*14
			}else if(cha4[i]=='F'){
				sum4=int(math.Pow(16,float64(len(cha4)-1-i)))*15
			}
		}else{
			sum4=int(math.Pow(16,float64(len(cha4)-1-i)))*inx
			
		}
		ans+=sum4
		
	}
	fmt.Println(ans)


	myslice:=[]string{"102","103","104"}
	var summ int = 0
	for i:=0;i<len(myslice);i++{
		sum1,err:=strconv.Atoi(myslice[i])
		if(err!=nil){
			fmt.Println("Error",err)
		}else{
			summ+=sum1
		}
	}
	fmt.Println(summ)

	var st3 string="10"
	f,err =strconv.ParseFloat(st3,64)
	if(err!=nil){
		fmt.Println("error",err)
	}else{
		fmt.Println(f*f)
	}
	var cha rune
	fmt.Println("Enter a character:")
	_, err = fmt.Scanf("%c", &cha) 
	if err != nil {
		fmt.Println("Invalid input. Please enter a single character.")
		return
	}

	m := int(cha) 
	fmt.Printf("The ASCII value of '%c' is: %d\n", cha, m)
}

// . Convert String Inputs to Integers and Calculate Sum:

// Write a program that takes five string inputs from the user, converts them to integers, and prints the sum.

// 2. Fahrenheit to Celsius Converter:

// Create a function that converts Fahrenheit to Celsius. Take temperature input from the user and display the result with two decimal places.

// 3. Integer to Float Conversion and Average Calculation:

// Take three integer inputs from the user, convert them to floats, and calculate the average using a function.

// 4. Boolean to Integer Conversion:

// Write a function that converts a boolean value to an integer (true -> 1, false -> 0). Take user input as "true" or "false" and convert it.

// 5. Convert Float to Integer and Check Even/Odd:

// Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd.

// 6. Hexadecimal to Decimal Converter:

// Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent.

// 7. Iterate Over a Slice of Strings and Convert to Integers:

// Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum.

// 8. String to Float Conversion with Error Handling:

// Take a string input representing a floating-point number, convert it to a float, and print the square of the number. Handle any conversion errors.

// 9. Read Multiple Values and Convert Types:

// Read an integer, a float, and a string from the user. Convert the integer to a float and the float to an integer, then display the converted values.

// 10. Character to ASCII Value Converter:

// Prompt the user to enter a single character, convert it to its ASCII integer value, and print the result.





