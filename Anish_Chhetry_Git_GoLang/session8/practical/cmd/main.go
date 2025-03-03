package main

import (
	"practical/internal/services"
)

func main() {
	data := services.NewDataProcessing()
	data.ReadJson("../jsonFiles/employees.json")
	//fmt.Println("Reading the JSON file: ", data)
	data.IncreaseSalary()
	//fmt.Println("\n\nReading the JSON file after increasing salary: ", data)
	data.SortEmployee()
	//fmt.Println("\n\nReading the JSON file after sorting: ", data)
	data.FilterEmp()
	data.WriteJson("../jsonFiles/updated_employees.json")
}
