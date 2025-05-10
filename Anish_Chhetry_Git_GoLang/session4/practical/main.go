package main

import (
	"practical/analyze"
)

func main() {
	bookings := []string{
		"Alice:FL123:Economy:120.50",
		"Bob:FL123:Business:450.00",
		"Charlie:FL456:Economy:150.75",
		"Alice:FL456:Economy:150.75",
		"Bob:FL123:Economy:120.50",
		"Anish:GH144:First:800.32",
		"Anish:GH145:First:800.33",
	}
	analyze.AnalyzeBookings(bookings)

}
