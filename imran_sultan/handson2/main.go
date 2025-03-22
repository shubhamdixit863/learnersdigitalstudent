package main

import (
	"fmt"
	"handson2/rental"
	"handson2/vehicle"
)

const name = "Rental Cost for Bike (9 hours): %.2f\n"

func main() {
	rm := rental.RentalManager{}
	rm.AddVehicle(&vehicle.Car{BaseVehicle: vehicle.BaseVehicle{"C123", 10}})
	rm.AddVehicle(&vehicle.Bike{BaseVehicle: vehicle.BaseVehicle{"B456", 5}})
	rm.AddVehicle(&vehicle.Truck{BaseVehicle: vehicle.BaseVehicle{"T789", 20}})

	rm.ListVehicles()

	rentalCost, err := rm.CalculateRental("C123", 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Rental Cost for Car (5 hours): %.2f\n", rentalCost)
	}

	NrentalCost, err := rm.CalculateRental("B456", -9)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf(name, NrentalCost)
	}
}
