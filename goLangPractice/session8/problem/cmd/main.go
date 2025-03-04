package main

import (
	"session8/problem/internal/model"
	empServices "session8/problem/internal/services"
)

func main() {
	empService := empServices.EmpService{
		Emp: []model.Employee{},
	}

	// Read File
	empService.Readfile("../data/employees.json")

	// Increase Salary
	empService.IncSalary()

	// Sorting Employees
	empService.SortEmployeesBySalary()

	// Filtering Employee
	empService.FilterEmployeesByDepartment("Engineering")

	// Summary Report
	empService.GenerateSummary()
}
