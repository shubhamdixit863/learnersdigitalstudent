package main

import (
	"fmt"
	"practical1/bookings"
)

func main() {
	bookings.Split()

	fmt.Println("\nTotal revenue for each flight:")
	m := bookings.TotalRevenue()
	for k, v := range m {
		fmt.Println(k, ":$", v)
	}

	fmt.Println("\nBookings by seat class:")
	m2 := bookings.SeatCount()
	for k, v := range m2 {
		fmt.Println(k, ":", v)
	}
	name, count := bookings.MostBooking()
	fmt.Println("\nPassenger with most booking:", name, count)
	flight, revenue := bookings.MostRevenueFlight()
	fmt.Printf("\nFlight with the highest revenue: %s ($%f)\n", flight, revenue)

}
