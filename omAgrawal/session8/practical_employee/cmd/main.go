package main

import (
	"practical_employee/internal/services"
)

func main() {

	empService := services.NewEmployeeService()
	//empService.AddEmployee(models.Employee{})
	empService.ReadJsonFile("./employee.json")
	empService.UpdateSalary()
	empService.FilterEmployee("Engineering")
	empService.ExportJson()

}
