package services

import (
	"encoding/json"
	"os"
	"revesion/internal/models"
)

type EmployeeServices struct {
	employee []models.Employee
}

func NewEmployeeServices(file string) (*EmployeeServices, error) {
	jsonfile, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// fmt.Println(jsonfile)
	var emp []models.Employee
	er1 := json.Unmarshal(jsonfile, &emp)
	if er1 != nil {
		return nil, er1
	}
	// fmt.Println(emp)
	return &EmployeeServices{
		employee: emp,
	}, nil
}
