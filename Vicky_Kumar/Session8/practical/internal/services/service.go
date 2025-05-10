package services

import (
	"encoding/json"
	"fmt"
	"os"
	"practical/internal/models"
	"sort"
)

type EmployeeService struct {
	Employees []models.Employee `json:"employees"`
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

func (e *EmployeeService) LoadEmployees(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err2 := json.Unmarshal(data, &e.Employees)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}

func (e *EmployeeService) IncreaseSalary() {
	for i := range e.Employees {
		if e.Employees[i].Age < 30 {
			e.Employees[i].Salary += 0.15 * e.Employees[i].Salary
		} else {
			e.Employees[i].Salary += 0.10 * e.Employees[i].Salary
		}
	}
}

func (e *EmployeeService) SortEmployees() {
	sort.Slice(e.Employees, func(i, j int) bool {
		return e.Employees[i].Salary > e.Employees[j].Salary
	})
}

func (e *EmployeeService) FilterEmployees(dept string) {
	for i, emp := range e.Employees {
		if emp.Department == dept {
			e.Employees = append(e.Employees[:i], e.Employees[i+1:]...)
		}
	}
}

func (e *EmployeeService) StoreEmployees(filepath string) {
	data, err := json.Marshal(e.Employees)
	if err != nil {
		fmt.Println(err)
		return
	}
	err2 := os.WriteFile(filepath, data, 0777)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}

func (e *EmployeeService) SummaryReport() {
	mp := map[string]int{}
	for _, emp := range e.Employees {
		mp[emp.Department]++
	}
	fmt.Println("Total Employees by Department:")
	for dep, n := range mp {
		fmt.Println(dep, ":", n)
	}
}
