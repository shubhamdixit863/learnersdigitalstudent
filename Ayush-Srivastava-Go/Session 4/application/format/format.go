package format

import (
	"fmt"
	"strconv"
	"strings"
)

var FlightToCostMap = make(map[string]float64)          // Revenue per flight
var SeatClassCount = make(map[string]int)               // Seat class count
var PassengerBooking = make(map[string]map[string]bool) // Passenger -> {Flight:Class -> true}

func GetMaps(bookings []string) {
	for _, booking := range bookings {
		parts := strings.Split(booking, ":")
		if len(parts) != 4 {
			continue 
		}

		passenger, flight, class, priceStr := parts[0], parts[1], parts[2], parts[3]

		cost, err := strconv.ParseFloat(priceStr, 64)
		if err != nil || cost < 0 {
			continue 
		}

		if _, exists := PassengerBooking[passenger]; !exists {
			PassengerBooking[passenger] = make(map[string]bool)
		}

		key := flight + ":" + class
		if PassengerBooking[passenger][key] {
			continue 
		}

		PassengerBooking[passenger][key] = true

		FlightToCostMap[flight] += cost

		SeatClassCount[class]++
	}
}

func GetTotalRevenueForEachFlight() {

	fmt.Println("\nRevenue per flight:")
	for flight, cost := range FlightToCostMap {
		fmt.Printf("%s: $%.2f\n", flight, cost)
	}
}

func GetSeatClassBookings() {
	fmt.Println("\nBookings by seat class:")
	for class, count := range SeatClassCount {
		fmt.Printf("%s: %d\n", class, count)
	}
}

func GetMostFrequentPassenger() {
	var mostFrequentPassenger string
	maxBookings := 0

	for passenger, flights := range PassengerBooking {
		if len(flights) > maxBookings {
			maxBookings = len(flights)
			mostFrequentPassenger = passenger
		}
	}

	fmt.Printf("\nPassenger with the most bookings: %s (%d bookings)\n", mostFrequentPassenger, maxBookings)
}

func GetFlightWithHighestRevenue() {
	var highestRevenueFlight string
	highestRevenue := 0.0

	for flight, revenue := range FlightToCostMap {
		if revenue > highestRevenue {
			highestRevenue = revenue
			highestRevenueFlight = flight
		}
	}

	fmt.Printf("Flight with the highest revenue: %s ($%.2f)\n", highestRevenueFlight, highestRevenue)
}

