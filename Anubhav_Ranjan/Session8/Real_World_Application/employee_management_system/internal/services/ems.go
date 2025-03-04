package services

import (
	"employee_management_system/internal/models"
	"employee_management_system/internal/utils"
	"encoding/json"
	"fmt"
)

type EmployeeManagementSystem struct {
	Employees []models.Employee
}

func NewEmployeeManagementSystem() *EmployeeManagementSystem {
	return &EmployeeManagementSystem{}
}

func (ems *EmployeeManagementSystem) LoadEmployees(filename string) error {
	data, err := utils.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &ems.Employees)
}

func (ems *EmployeeManagementSystem) SaveEmployees(filename string) error {
	data, err := json.MarshalIndent(ems.Employees, "", "  ")
	if err != nil {
		return err
	}
	return utils.WriteFile(filename, data)
}

func (ems *EmployeeManagementSystem) UpdateSalaries() {
	for i := range ems.Employees {
		if ems.Employees[i].Age < 30 {
			ems.Employees[i].Salary *= 1.15
		} else if ems.Employees[i].Age > 50 {
			ems.Employees[i].Salary *= 1.10
		}
	}
}

func (ems *EmployeeManagementSystem) SortBySalary() {
	n := len(ems.Employees)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if ems.Employees[j].Salary < ems.Employees[j+1].Salary {
				ems.Employees[j], ems.Employees[j+1] = ems.Employees[j+1], ems.Employees[j]
			}
		}
	}
}

func (ems *EmployeeManagementSystem) FilterByDepartment(department string) []models.Employee {
	var filteredEmployees []models.Employee
	for _, e := range ems.Employees {
		if e.Department == department {
			filteredEmployees = append(filteredEmployees, e)
		}
	}
	return filteredEmployees
}

func (ems *EmployeeManagementSystem) PrintSummary() {
	deptCount := make(map[string]int)
	for _, e := range ems.Employees {
		deptCount[e.Department]++
	}
	fmt.Println("Total Employees by Department:")
	for dept, count := range deptCount {
		fmt.Printf("%s: %d\n", dept, count)
	}
}
