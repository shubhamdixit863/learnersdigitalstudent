package main

import (
	"flight_Booking_Analysis_System/analysis"
	"flight_Booking_Analysis_System/database"
	"fmt"
)

func main() {
	bookings := database.GetBookings()

	flightRevenue := analysis.CalculateFlightRevenue(bookings)
	seatClassCount := analysis.CountSeatClassBookings(bookings)
	passenger, maxBookings := analysis.FindMostRegularPassenger(bookings)
	highestRevFlight, highestRevenue := analysis.FindHighestRevenueFlight(flightRevenue)

	fmt.Println("\n Flight Booking Analysis Report: ")
	fmt.Println("--------------------------------")

	fmt.Println("\n1. Revenue per flight: ")
	for flight, revenue := range flightRevenue {
		fmt.Printf("   %s: $%.2f\n", flight, revenue)
	}

	fmt.Println("\n2. Bookings by Seat Class:")
	for _, class := range []string{"Economy", "Business", "First"} {
		fmt.Printf("   %s: %d\n", class, seatClassCount[class])
	}

	// 3. Most frequent flyer
	fmt.Printf("\n3. Passenger with the most bookings: %s (%d bookings)\n", passenger, maxBookings)

	// 4. Highest revenue flight
	fmt.Printf("\n4. Flight with Highest Revenue: %s ($%.2f)\n\n", highestRevFlight, highestRevenue)

}
