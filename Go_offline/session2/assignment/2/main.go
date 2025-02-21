package main

import "fmt"

func main() {

	var a float64
	fmt.Scanln(&a)
	fmt.Printf("The area of circle with radius %f is %f\n", a, area(a))
}
 func area(radius float64) float64 {
	return 3.14 * radius * radius
}
