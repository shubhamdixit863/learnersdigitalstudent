package main

import "revesion/internal/services"

func main() {
	services.NewEmployeeServices("../employee.json")
}
