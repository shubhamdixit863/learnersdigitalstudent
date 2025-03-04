package main

import (
	"employee_management_system/internal/services"
	"fmt"
)

func main() {
	ems := services.NewEmployeeManagementSystem()

	err := ems.LoadEmployees("../data/employees.json")
	if err != nil {
		fmt.Printf("Error loading employees: %v", err)
	}

	ems.UpdateSalaries()
	ems.SortBySalary()

	err = ems.SaveEmployees("../data/updated_employees.json")
	if err != nil {
		fmt.Printf("Error saving employees: %v", err)
	}

	ems.PrintSummary()

	dept := "Engineering"
	filtered := ems.FilterByDepartment(dept)
	fmt.Printf("\nEmployees in %s:\n", dept)
	for _, e := range filtered {
		fmt.Println(e)
	}
}
