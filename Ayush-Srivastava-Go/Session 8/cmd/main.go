package main

import (
	"fmt"
	"session8/internal/services"
)

func main() {
	employeeService, err := services.NewEmployeeService("internal/data/employees.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("Employees:")
	fmt.Println(employeeService)

	employeeService.IncreaseSalary()
	employeeService.SortbySalary()

	err = employeeService.MarshalData("internal/data/updated_employees.json")
	if err != nil {
		panic(err)
	}

	engineeringEmployees := employeeService.FilterByDepartment("Engineering")
	financeEmployees := employeeService.FilterByDepartment("Finance")

	fmt.Println("Engineering Employees:")
	fmt.Println(engineeringEmployees)

	fmt.Println("Finance Employees:")
	fmt.Println(financeEmployees)

	summary := employeeService.GetSummary()
	fmt.Println("Total Employees by Department:")
	for dept, count := range summary {
		fmt.Printf("%s: %d\n", dept, count)
	}
}
