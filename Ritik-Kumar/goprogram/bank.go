package main

import (
	"fmt"
	
)

// func fizzbuz(s int) {
// 	if s%5 == 0 {
// 		fmt.Println("fizz")
// 	} else if s%3 == 0 {
// 		fmt.Println("BUZZ")
// 	} else {
// 		fmt.Println("none")
// 	}
// }



func main() {
	var a int
	var b int
	var step string
	var c int
	
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&step)

	if(step=="add"){
		c=a+b;
	}else if(step=="minus"){
		c=a-b;
	}else if(step=="multiply"){
		c=a*b;
	}else if(step=="divide"){
		c=a/b;
	}

	

	fmt.Println(c)
}


//you have to write a program to take the user details as input (name,age,address,company,parents name)
//convert the age into days
//print out the whole details