package vehicle

type Truck struct {
	BaseVehicle
}

func (t *Truck) GetVehicleType() string {
	return "Truck"
}
