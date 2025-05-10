package services

type Truck struct {
	model      string
	hourlyRate float64
}

func NewTruck(model string, hourlyRate float64) *Truck {
	return &Truck{model: model, hourlyRate: hourlyRate}
}

func (t *Truck) GetType() string {
	return "Truck"
}

func (t *Truck) GetHourlyRate() float64 {
	return t.hourlyRate
}

func (t *Truck) GetModel() string {
	return t.model
}