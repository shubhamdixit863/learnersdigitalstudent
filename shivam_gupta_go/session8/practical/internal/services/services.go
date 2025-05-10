package services

import (
	"encoding/json"
	"fmt"
	"os"
	"practical/internal/model"
	"sort"
)

type Operations struct {
	Data []model.Person
}

func NewOperations()*Operations{
	return &Operations{}
}

func (op *Operations) ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (op *Operations) Unmarshal(data []byte) ([]model.Person, error) {
	err := json.Unmarshal(data, &op.Data)
	if err != nil {
		fmt.Println(err)
	return nil, err
	}
	return op.Data, err
}

func (op *Operations) SalaryIncr()  {
	for i := range op.Data {
		if op.Data[i].Age < 30 {
			op.Data[i].Salary += op.Data[i].Salary * 0.15
		} else if op.Data[i].Age > 50 {
			op.Data[i].Salary += op.Data[i].Salary * 0.10
		}
	}
}

func (op *Operations) SortEmployeesBySalary() {
	sort.Slice(op.Data, func(i, j int) bool {
		return op.Data[i].Salary > op.Data[j].Salary // Descending order
	})
}

func (op *Operations) FilterEmp() {
	fmt.Println("Enter department name:")
	var dep string
	fmt.Scanln(&dep)
	for _, emp := range op.Data {
		if emp.Department == dep {
			fmt.Println(emp)
		}
	}
}

func (op *Operations) CreateFile() {
	file, err := os.Create("../Mydata.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	b, err := json.Marshal(op.Data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	_, err = file.Write(b)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func (op *Operations) Summary() {
	fmt.Println("Total employees by department")
	deptCount := make(map[string]int)
	for _, emp := range op.Data {
		deptCount[emp.Department]++
	}
	for key, value := range deptCount {
		fmt.Printf("%s: %d\n", key, value)
	}
}
