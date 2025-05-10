package services

type Vehicle interface {
	GetType() string
	GetHourlyRate() float64
	GetModel() string
}