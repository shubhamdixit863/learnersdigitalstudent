package vehicle

type Bike struct {
	BaseVehicle
}

func (b *Bike) GetVehicleType() string {
	return "Bike"
}
