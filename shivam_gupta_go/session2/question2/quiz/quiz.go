package quiz

import "fmt"
var Score int
func Quiz() int {
	
	Score = 0
	fmt.Println("Q1.What is the capital of France? ")
	fmt.Println("1) Berlin   2) Madrid 3) Paris  4) Rome")
	var a int
	fmt.Scanln(&a)
	if a == 3 {
		fmt.Println("corret")
		Score++
	} else {
		fmt.Println("wrong")
	}

	fmt.Println("Q2.What is 5 + 7? ")
	fmt.Println("1) 12   2) 23 3) 15  4) 8")
	var b int
	fmt.Scanln(&b)
	if b == 1 {
		fmt.Println("corret")
		Score++
	} else {
		fmt.Println("wrong")
	}

	fmt.Println("Q3.Which language is used for web apps? ")
	fmt.Println("1) Python 2) JavaScript 3) C++  4) Go")
	var c int
	fmt.Scanln(&c)
	if c == 2 {
		fmt.Println("corret")
		Score++
	} else {
		fmt.Println("wrong")
	}

	fmt.Println("Q4.What is the capital of Japan? ")
	fmt.Println("1) delhi   2) tokyo 3) paris  4) Rome")
	var d int
	fmt.Scanln(&d)
	if d == 2 {
		fmt.Println("corret")
		Score++
	} else {
		fmt.Println("wrong")

	}
	fmt.Println("Q5.How many continents are there?")
	fmt.Println("1) 3   2) 5  3) 7  4) 8")
	var e int
	fmt.Scanln(&e)
	if e == 3 {
		fmt.Println("corret")
		Score++
	} else {
		fmt.Println("wrong")
	}
  
	fmt.Printf("your score %d\n",Score)
	return Score
}


