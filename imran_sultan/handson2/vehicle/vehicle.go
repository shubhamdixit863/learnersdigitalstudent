package vehicle

import "errors"

type Vehicle interface {
	GetVehicleType() string
	GetVehicleID() string
	CalculateRentalCost(hours int) (float64, error)
	GetRatePerHour() float64
}

type BaseVehicle struct {
	VehicleID   string
	RatePerHour float64
}

func (v *BaseVehicle) GetVehicleID() string {
	return v.VehicleID
}

func (v *BaseVehicle) GetRatePerHour() float64 {
	return v.RatePerHour
}

func (v *BaseVehicle) CalculateRentalCost(hours int) (float64, error) {
	if hours <= 0 {
		return 0, errors.New("rental hours must be greater than zero")
	}
	return v.RatePerHour * float64(hours), nil
}
