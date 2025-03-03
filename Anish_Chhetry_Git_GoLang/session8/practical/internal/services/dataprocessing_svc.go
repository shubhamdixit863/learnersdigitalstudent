package services

import (
	"encoding/json"
	"fmt"
	"os"
	"practical/internal/model"
)

type DataProcessing struct {
	Employees []model.Employee
}

func NewDataProcessing() *DataProcessing {
	return &DataProcessing{}
}

func (d *DataProcessing) IncreaseSalary() {
	for i := range d.Employees {
		if d.Employees[i].Age < 30 {
			d.Employees[i].Salary += d.Employees[i].Salary * 0.15
		} else if d.Employees[i].Age > 50 {
			d.Employees[i].Salary += d.Employees[i].Salary * 0.1
		}
	}

}

func (d *DataProcessing) SortEmployee() {
	for i := 0; i < len(d.Employees)-1; i++ {
		max := i
		for j := i + 1; j < len(d.Employees); j++ {
			if d.Employees[j].Salary > d.Employees[max].Salary {
				max = j
			}
		}
		d.Employees[i], d.Employees[max] = d.Employees[max], d.Employees[i]
	}
}
func (d *DataProcessing) FilterEmp() {
	departmentMap := make(map[string][]model.Employee)
	fmt.Println("Total Employees by Department:")
	for _, emp := range d.Employees {
		departmentMap[emp.Department] = append(departmentMap[emp.Department], emp)
	}
	for dept := range departmentMap {
		fmt.Print(dept, ":")
		for _, val := range d.Employees {
			count := 0
			if val.Department == dept {
				count++
				fmt.Println(count)
			}
		}
	}

}

func (d *DataProcessing) ReadJson(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, &d.Employees)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}
}
func (d *DataProcessing) WriteJson(filename string) {
	s := d
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	err1 := os.WriteFile(filename, data, 0644)
	if err1 != nil {
		fmt.Println(err)
	}
}
