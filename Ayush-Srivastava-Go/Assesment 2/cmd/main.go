package main

import (
	"fmt"
	"log"
	"vehicle-rent/internal/services"
)

func main() {
	rentalService := services.NewRentalService()
	
	car1 := services.NewCar("Toyota", 10)
	car2 := services.NewCar("Honda", 15)

	bike1 := services.NewBike("Yamaha", 5)
	bike2 := services.NewBike("Suzuki", 3)

	truck1 := services.NewTruck("Isuzu", 20)
	truck2 := services.NewTruck("Hino", 25)

	nonameError := services.NewCar("", 0)
	invalidHoursError := services.NewBike("Kawasaki", 0)



	vehicles := []services.Vehicle{car1, car2, bike1, bike2, truck1, truck2, nonameError, invalidHoursError}

	for _, vehicle := range vehicles {
		cost, err := rentalService.CalculateRentalCost(vehicle, 5)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("Vehicle: %s | Model: %s | Cost: $%.2f\n", vehicle.GetType(), vehicle.GetModel(), cost)
	}

}