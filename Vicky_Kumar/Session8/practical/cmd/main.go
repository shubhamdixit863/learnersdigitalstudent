package main

import (
	"practical/internal/services"
)

func main() {
	emp_service := services.NewEmployeeService()

	emp_service.LoadEmployees("./data/employees.json")
	emp_service.IncreaseSalary()
	emp_service.SortEmployees()
	emp_service.FilterEmployees("Engineering")
	emp_service.StoreEmployees("./data/updated_employees.json")
	emp_service.SummaryReport()
}
