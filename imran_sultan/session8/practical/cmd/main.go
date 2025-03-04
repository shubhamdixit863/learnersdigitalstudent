package main

import (
	"fmt"
	"practical/internal/services"
)

func main() {
	obj := services.NewEmployeeServices()

	employeeService, err := obj.Readjsonfile("../employee.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Employees:")
	fmt.Println(employeeService)
	employeeService.IncreaseSalary()
	employeeService.SortbySalary()
	fmt.Println("after Employees:")
	fmt.Println(employeeService)

}
