package empServices

import (
	"encoding/json"
	"fmt"
	"os"
	"session8/problem/internal/model"
	"sort"
)

type EmpService struct {
	Emp []model.Employee
}

func (emp *EmpService) Readfile(filename string) {
	data, e := os.ReadFile(filename)

	if e != nil {
		return
	}

	var employees []model.Employee

	s := []byte(data)
	err := json.Unmarshal(s, &employees)
	if err != nil {
		fmt.Println(err)
	}

	emp.Emp = employees

	fmt.Println(*emp)
}

func (emp *EmpService) Writefile(filename string) {
	data, e := json.Marshal(emp)
	if e != nil {
		fmt.Println("Error marshaling JSON:", e)
		return
	}

	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully!")
}

func (emp *EmpService) IncSalary() {
	for i, e := range emp.Emp {
		if e.Age < 30 {
			emp.Emp[i].Salary *= 1.15
		} else if e.Age > 50 {
			emp.Emp[i].Salary *= 1.10
		}
	}

	emp.Writefile("../data/updated_employee.json")
}

func (emp *EmpService) SortEmployeesBySalary() {
	sort.Slice(emp.Emp, func(i, j int) bool {
		return emp.Emp[i].Salary > emp.Emp[j].Salary
	})

	emp.Writefile("../data/updated_employee.json")
}

func (emp *EmpService) FilterEmployeesByDepartment(department string) {
	var filtered []model.Employee
	for _, employee := range emp.Emp {
		if employee.Department == department {
			filtered = append(filtered, employee)
		}
	}
	fmt.Println(filtered)
}

func (emp *EmpService) GenerateSummary() {
	departmentCount := make(map[string]int)
	for _, employee := range emp.Emp {
		departmentCount[employee.Department]++
	}

	fmt.Println("Total Employees by Department:")
	for department, count := range departmentCount {
		fmt.Printf("%s: %d\n", department, count)
	}
}
