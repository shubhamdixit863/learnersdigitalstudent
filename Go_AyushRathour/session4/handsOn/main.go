package main

import (
	"fmt"
	"handsOn/dataUpdate"
	"handsOn/maxbooking"
	"handsOn/maxrevenue"
)

func main() {

	bookings := []string{
		"Alice:FL123:Economy:120.50",
		"Bob:FL123:Business:450.00",
		"Charlie:FL456:Economy:150.75",
		"Alice:FL456:Economy:150.75",
		"Bob:FL123:Economy:120.50",
	}

	name_map, flightName_map, seatclass_map := dataUpdate.AddData(bookings)

	fmt.Printf("Revenue per flight: ")
	for key, value := range flightName_map {
		fmt.Printf("%s: %.2f\n", key, value)
	}

	fmt.Printf("Number of bookings in each seat class: ")
	for key, value := range seatclass_map {
		fmt.Printf("%s: %d\n", key, value)
	}

	max_name := maxbooking.MaxBooking(name_map)
	fmt.Printf("Passenger with most bookings: %s\n", max_name)

	max_flight := maxrevenue.MaxRevenue(flightName_map)
	fmt.Printf("Flight with most revenue: %s\n", max_flight)

}
