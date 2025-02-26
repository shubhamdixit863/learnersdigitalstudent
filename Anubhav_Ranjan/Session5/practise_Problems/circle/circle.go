package circle

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{Radius: radius}
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}
