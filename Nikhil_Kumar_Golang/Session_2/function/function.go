package main

import "fmt"

func square(num int) int {
	return num*num
}

// multiple return values
func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero %s", "divided by zero")
	}
	g:= num1/num2
	return g, nil	
	
}

// Variadic Functions
func variadicFunc(f float64, n ...int) {
	fmt.Println(f, n)
	for i:=0; i<len(n); i++ {
		fmt.Println(n[i])
	}
}


// named return

func namedReturn(n int) (double int)  {
	double = n * 2
	return //explicitly return
}



func main() {
	// square(3)
	fmt.Println(square(3))
	fmt.Println(divide(3, 0))

	// to store to value from function
	a,b := divide(8, 2)
	fmt.Println(a,b)

	// handling error 
	a1,err := divide(8, 0)
	fmt.Println(a1,err)

	if err != nil {
		fmt.Println(err)
	}
	// Variadic Functions
	fmt.Println("Variadic Functions")
	variadicFunc(1.0, 2, 3, 4, 5, 6, 7, 8, 9)
}