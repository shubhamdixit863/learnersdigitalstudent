package rectangle

type Rectangle struct {
	Length  float64
	Breadth float64
}

func NewRectangle(length float64, breadth float64) Rectangle {
	return Rectangle{Length: length, Breadth: breadth}
}
func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}
