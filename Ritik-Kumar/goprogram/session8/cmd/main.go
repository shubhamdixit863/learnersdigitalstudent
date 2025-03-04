package main

import (
	"fmt"
	"session8/internal/services"
)

func main() {
	service := services.EmployeeService{}

	err := service.ReadEmployeesFromFile("internal/data/employees.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	service.UpdateSalaries()

	service.SortEmployeesBySalary()

	service.GenerateSummaryReport()

	fmt.Println("\nEnter department to filter employees (or press Enter to skip):")
	var dept string
	fmt.Scanln(&dept)
	if dept != "" {
		service.FilterEmployeesByDepartment(dept)
	}
}
