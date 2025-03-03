package services

import (
	"encoding/json"
	"fmt"
	"os"
	"practical_employee/internal/models"
	"sort"
)

type EmployeeService struct {
	Employees []*models.Employee `json:"employees"`
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

func (e *EmployeeService) ReadJsonFile(path string) {
	//[]models.Employee
	data, _ := os.ReadFile("./employee.json")
	fmt.Println(string(data))
	// employee := Employees{}
	eerr := json.Unmarshal(data, &e.Employees)

	if eerr != nil {
		//fmt.Println(data)
		fmt.Println(eerr)
	}
	// fmt.Println(employee)
	for _, d := range e.Employees {
		fmt.Println(*d)
	}

}

func (e *EmployeeService) UpdateSalary() {
	for _, d := range e.Employees {
		if d.Age < 30 {
			d.Salary += d.Salary * .15
		}

		if d.Age >= 50 {
			d.Salary -= d.Salary * .10

		}
	}

	for _, d := range e.Employees {
		fmt.Println(*d)
	}
}

func (e *EmployeeService) SortEmployee() {
	sort.Slice(e.Employees, func(i, j int) bool {
		return e.Employees[i].Salary < e.Employees[j].Salary
	})
	fmt.Println(e.Employees)
}

func (e *EmployeeService) FilterEmployee(dep string) {

	for _, d := range e.Employees {

		if d.Department == dep {
			fmt.Println(*d)
		}
	}
}

func (e *EmployeeService) ExportJson() {
	data, _ := json.Marshal(e)

	err := os.WriteFile("UpdatedEmployee.json", data, 0777)

	if err != nil {
		fmt.Println(err)
	}

}
