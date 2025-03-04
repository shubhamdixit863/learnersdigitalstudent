package services

import (
	"encoding/json"
	"session8/internal/models"
	"fmt"
	"os"
	"sort"
)

type EmployeeService struct {
	Employees []models.Employee
}

func (es *EmployeeService) ReadEmployeesFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &es.Employees)
	if err != nil {
		return err
	}

	return nil
}

func (es *EmployeeService) UpdateSalaries() {
	for i := range es.Employees {
		es.Employees[i].UpdateSalary()
	}
}

func (es *EmployeeService) SortEmployeesBySalary() {
	sort.Slice(es.Employees, func(i, j int) bool {
		return es.Employees[i].Salary > es.Employees[j].Salary
	})
}

func (es *EmployeeService) GenerateSummaryReport() {
	deptCount := make(map[string]int)
	for _, emp := range es.Employees {
		deptCount[emp.Department]++
	}

	fmt.Println("\nTotal Employees by Department:")
	for dept, count := range deptCount {
		fmt.Printf("%s: %d\n", dept, count)
	}
}

func (es *EmployeeService) FilterEmployeesByDepartment(department string) {
	fmt.Printf("\nEmployees in %s Department:\n", department)
	for _, emp := range es.Employees {
		if emp.Department == department {
			fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", emp.ID, emp.Name, emp.Salary)
		}
	}
}
