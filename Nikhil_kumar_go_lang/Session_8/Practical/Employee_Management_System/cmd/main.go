package main

import (
	"Employee_Management_System/model"
	"Employee_Management_System/services"
	"fmt"
	"os"
)

func main() {
	fmt.Println("***********Employee Management System*********************")

	empService := services.EmployeeService{
		Employees: []model.Employee{},
	}

	directory, err := os.Getwd() //get the current directory using the built-in function
	if err != nil {
		fmt.Println(err) //print the error if obtained
	}
	fmt.Println("Current working directory:", directory) //p

	fileName := "../data/employees.json"
	employees, err := empService.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(employees) > 0 {
		empService = employees[0]
	}

	fmt.Println("Employee data loaded successfully")

	empService.IncresedSalary()
	fmt.Println("Salaries increased successfully", empService.Employees)

	empService.Sort_by_Salary()
	fmt.Println("Employees sorted by salary", empService.Employees)

	department := "Engineering"

	filteredEmployees, err := empService.Filter_EmployeeService_by_department(department)

	if err != nil {
		fmt.Println("Error filtering employees:", err)
		return
	}

	fmt.Printf("Filtered %d employees in %s department\n", len(filteredEmployees), department)

	err = empService.WriteFile("data/updated_employees.json", []services.EmployeeService{empService})
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Updated employee data written to updated_employees.json")

	fmt.Println("\nSummary report:")
	summaryReport := empService.Summary_report()
	for dept, count := range summaryReport {
		fmt.Printf("%s = %d\n", dept, count)
	}
}
