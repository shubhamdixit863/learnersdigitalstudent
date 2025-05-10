package services

type Bike struct {
	model      string
	hourlyRate float64
}

func NewBike(model string, hourlyRate float64) *Bike {
	return &Bike{model: model, hourlyRate: hourlyRate}
}

func (b *Bike) GetType() string {
	return "Bike"
}

func (b *Bike) GetHourlyRate() float64 {
	return b.hourlyRate
}

func (b *Bike) GetModel() string {
	return b.model
}