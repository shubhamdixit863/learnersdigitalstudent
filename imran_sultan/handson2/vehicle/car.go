package vehicle

type Car struct {
	BaseVehicle
}

func (c *Car) GetVehicleType() string {
	return "Car"
}
