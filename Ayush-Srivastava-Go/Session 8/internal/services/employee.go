package services

import (
	"encoding/json"
	"os"
	"session8/internal/models"
)

type EmployeeService struct {
	employees []models.Employee
}

func (es *EmployeeService) IncreaseSalary() {

	for i := range es.employees {
		if es.employees[i].Age < 30 {
			es.employees[i].Salary += (15.0 / 100.0) * es.employees[i].Salary
		} else if es.employees[i].Age >= 30 && es.employees[i].Age < 50 {
			continue
		} else {
			es.employees[i].Salary += (10.0 / 100.0) * es.employees[i].Salary
		}
	}
}

func (es *EmployeeService) SortbySalary() {

	for i := range es.employees {
		for j := range es.employees {
			if es.employees[i].Salary > es.employees[j].Salary {
				es.employees[i], es.employees[j] = es.employees[j], es.employees[i]
			}
		}
	}
}

func (es *EmployeeService) FilterByDepartment(department string) []models.Employee {
	filteredEmployees := []models.Employee{}
	for _, employee := range es.employees {
		if employee.Department == department {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}
	return filteredEmployees
}



func (es *EmployeeService) GetSummary() map[string]int {
	summary := make(map[string]int)
	for _, employee := range es.employees {
		summary[employee.Department]++
	}
	return summary
}


func NewEmployeeService(filename string) (*EmployeeService, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var employees []models.Employee

	err = json.Unmarshal(data, &employees)

	if err != nil {
		return nil, err
	}

	return &EmployeeService{employees: employees}, nil
}

func (es *EmployeeService) MarshalData(filename string) error {
	data, err := json.Marshal(es.employees)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

