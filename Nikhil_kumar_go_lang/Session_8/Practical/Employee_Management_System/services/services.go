package services

import (
	"Employee_Management_System/model"
	"encoding/json"
	"os"
)

type EmployeeService struct {
	Employees []model.Employee
}

func (emp *EmployeeService) ReadFile(fileName string) ([]EmployeeService, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var employees []model.Employee
	err = json.Unmarshal(data, &employees)
	if err != nil {
		return nil, err
	}
	emp.Employees = employees
	return []EmployeeService{*emp}, nil
}

func (emp *EmployeeService) IncresedSalary() {
	for i := range emp.Employees {
		if emp.Employees[i].Age < 30 {
			emp.Employees[i].Salary *= 1.15
		}
		if emp.Employees[i].Age > 50 {
			emp.Employees[i].Salary *= 1.10
		}
	}
}

func (emp *EmployeeService) Sort_by_Salary() {
	for i := range emp.Employees {
		for j := i + 1; j < len(emp.Employees); j++ {
			if emp.Employees[i].Salary < emp.Employees[j].Salary {
				emp.Employees[i], emp.Employees[j] = emp.Employees[j], emp.Employees[i]
			}
		}
	}
}

func (emp *EmployeeService) Filter_EmployeeService_by_department(department string) ([]EmployeeService, error) {
	var filteredEmp []EmployeeService
	for i := range emp.Employees {
		if emp.Employees[i].Department == department {
			Employees := EmployeeService{Employees: []model.Employee{emp.Employees[i]}}
			filteredEmp = append(filteredEmp, Employees)

		}
	}

	return filteredEmp, nil
}

func (emp *EmployeeService) WriteFile(fileName string, data []EmployeeService) error {
	emp_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, emp_data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (emp *EmployeeService) Summary_report() map[string]int {
	emp_count := make(map[string]int)
	for i := range emp.Employees {
		emp_count[emp.Employees[i].Department]++
	}
	return emp_count
}
