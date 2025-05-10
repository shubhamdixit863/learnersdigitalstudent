package services

import "vehicle-rent/internal/errors"

type RentalService struct{}

func NewRentalService() *RentalService {
	return &RentalService{}
}

func (rs *RentalService) CalculateRentalCost(vehicle Vehicle, hours int) (float64, error) {

	if vehicle == nil {
		return 0, errors.ErrNilVehicle
	}
	if hours <= 0 {
		return 0, errors.ErrInvalidHours
	}
	if vehicle.GetModel() == "" {
		return 0, errors.ErrUnNamedVehicle
	}
	if vehicle.GetHourlyRate() <= 0 {
		return 0, errors.ErrInvalidRentalRate
	}
	return float64(hours) * vehicle.GetHourlyRate(), nil
}