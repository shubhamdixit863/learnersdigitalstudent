package services

type Car struct {
	model      string
	hourlyRate float64
}

func NewCar(model string, hourlyRate float64) *Car {
	return &Car{model: model, hourlyRate: hourlyRate}
}

func (c *Car) GetType() string {
	return "Car"
}

func (c *Car) GetHourlyRate() float64 {
	return c.hourlyRate
}

func (c *Car) GetModel() string {
	return c.model
}