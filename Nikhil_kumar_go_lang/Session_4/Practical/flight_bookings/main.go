package main

import (
	"flight_booking/display"
)

var Booking []string = []string{
	"Alice:FL123:Economy:120.50",
	"Bob:FL123:Business:450.00",
	"Charlie:FL456:Economy:150.75",
	"Alice:FL456:Economy:150.75",
	"Bob:FL123:Economy:120.50",
}

func main() {
	display.Display(Booking)

}
