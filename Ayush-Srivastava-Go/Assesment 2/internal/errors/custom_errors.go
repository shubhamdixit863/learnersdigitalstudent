package errors

import "errors"

var (
	ErrInvalidHours = errors.New("rental hours must be greater than zero")
	ErrNilVehicle   = errors.New("vehicle cannot be nil")
	ErrUnNamedVehicle = errors.New("vehicle must have a name")
	ErrInvalidRentalRate = errors.New("rental rate must be greater than zero")
)
