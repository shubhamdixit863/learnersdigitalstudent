package main

import (
	"fmt"
	"practical/internal/services"
)

func main() {
    op:=services.NewOperations()

	data, err := op.ReadFile("employee.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	//  fmt.Println(data)

	data2, _ := op.Unmarshal(data)

	fmt.Println(data2)

	op.SalaryIncr()

	op.SortEmployeesBySalary()

	op.FilterEmp()

	op.CreateFile()

	op.Summary()
}
