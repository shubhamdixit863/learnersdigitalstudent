package rental

import (
	"errors"
	"fmt"
	"handson2/vehicle"
)

type RentalManager struct {
	Vehicles []vehicle.Vehicle
}

func (rm *RentalManager) AddVehicle(v vehicle.Vehicle) {
	rm.Vehicles = append(rm.Vehicles, v)
}

func (rm *RentalManager) CalculateRental(vehicleID string, hours int) (float64, error) {
	if hours <= 0 {
		return 0, errors.New("rental hours must be positive")
	}

	for _, vehicle := range rm.Vehicles {
		if vehicle.GetVehicleID() == vehicleID {
			return vehicle.CalculateRentalCost(hours)
		}
	}
	return 0, errors.New("vehicle not found")
}

func (rm *RentalManager) ListVehicles() {
	for _, vehicle := range rm.Vehicles {
		fmt.Printf("%s (ID: %s) - Rate: %.2f/hour\n", vehicle.GetVehicleType(), vehicle.GetVehicleID(), vehicle.GetRatePerHour())
	}
}
