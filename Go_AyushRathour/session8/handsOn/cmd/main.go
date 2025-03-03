package main

import (
	"fmt"
	"handsOn/internal/services"
)

func main() {

	emp := services.NewEmployee()

	err := emp.ReadFile("employees.json")
	if err != nil {
		fmt.Println(err)
	}
	emp.IncreaseSalary()
	emp.Sort()
	emp.FilterEmployees("marketing")
	emp.WriteUpdatedData()

	fmt.Println("Summary Report")

	emp.PrintSummary()

}
