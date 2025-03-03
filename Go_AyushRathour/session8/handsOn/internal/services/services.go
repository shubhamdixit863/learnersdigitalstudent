package services

import (
	"encoding/json"
	"fmt"
	"handsOn/internal/model"
	"os"
	"sort"
)

type Model struct {
	Employee []model.Employee
}

func NewEmployee() *Model {
	return &Model{}
}

func (emp *Model) ReadFile(filename string) error {
	byteValue, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err2 := json.Unmarshal(byteValue, &emp.Employee)
	return err2
}

func (emp *Model) IncreaseSalary() {
	for i := range emp.Employee {
		if emp.Employee[i].Age < 30 {
			emp.Employee[i].Salary *= 1.15
		} else if emp.Employee[i].Age > 50 {
			emp.Employee[i].Salary *= 1.10
		}
	}

}

func (emp *Model) Sort() {

	sort.Slice(emp.Employee, func(i, j int) bool {
		return emp.Employee[i].Salary > emp.Employee[j].Salary
	})
}

func (emp *Model) FilterEmployees(dept string) {
	for _, emp := range emp.Employee {
		if emp.Department == dept {
			fmt.Printf("ID: %d, Name: %s, Age: %d, Department: %s, Salary: %.2f\n", emp.ID, emp.Name, emp.Age, emp.Department, emp.Salary)
			fmt.Printf("Address: %s, %s - %s\n", emp.Address.Street, emp.Address.City, emp.Address.ZipCode)
			fmt.Printf("Skills: %v\n", emp.Skills)
		}

	}

}

func (emp *Model) WriteUpdatedData() {
	updatedData, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling updated data: %v", err)
	}

	err = os.WriteFile("updated_employees.json", updatedData, 0644)
	if err != nil {
		fmt.Printf("Failed to write updated JSON file: %v", err)
	}

	fmt.Println("\n Updated employee data saved to 'updated_employees.json'.")
}

func (emp *Model) PrintSummary() {
	departmentCount := make(map[string]int)
	for _, emp := range emp.Employee {
		departmentCount[emp.Department]++
	}

	fmt.Println("\n Employee Count per Department:")
	for dept, count := range departmentCount {
		fmt.Printf("%s: %d employees\n", dept, count)
	}
}
