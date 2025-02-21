// Create a function that checks if a given year is a leap year.


package main

import "fmt"

func main() {
	var year int
	fmt.Scanln(&year)
	fmt.Printf("The year is a leap year: %t\n", isLeapYear(year))
}

func isLeapYear(year int) bool {
	if year % 4 == 0 {
		if year % 100 == 0 {
			if year % 400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

