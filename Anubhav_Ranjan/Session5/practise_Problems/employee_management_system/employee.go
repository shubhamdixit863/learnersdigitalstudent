package employee_management_system

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func AddEmployee(db *[]Employee, id int, name string, salary float64) {
	newEmployee := Employee{ID: id, Name: name, Salary: salary}
	*db = append(*db, newEmployee)
	fmt.Printf(" Employee Added: %+v\n", newEmployee)
}

func RemoveEmployee(db *[]Employee, id int) error {
	for i, emp := range *db {
		if emp.ID == id {
			*db = append((*db)[:i], (*db)[i+1:]...)
			fmt.Printf(" Employee with ID %d is Removed..\n", id)
			return nil
		}
	}

	return errors.New("Employee NOT Found")
}

func GiveSalaryHike(db *[]Employee, threshold, hikePercent float64) {
	for _, emp := range *db {
		if emp.Salary < threshold {
			newSalary := emp.Salary * (1 + hikePercent/100)
			fmt.Printf("Salary of %s Hiked from %.2f to %.2f\n", emp.Name, emp.Salary, newSalary)
			emp.Salary = newSalary
		}
	}
}

func DisplayEmployees(db *[]Employee) {
	if len(*db) == 0 {
		fmt.Println("No Employees Found")
		return
	}

	fmt.Println("\n Employee Records:")
	for _, emp := range *db {
		fmt.Printf("ID : %d | Name : %s | Salary : %.2f\n", emp.ID, emp.Name, emp.Salary)
	}
}
