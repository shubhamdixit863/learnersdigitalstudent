package analysis

import "strconv"

func CalculateFlightRevenue(bookings [][]string) map[string]float64 {
	flightRevenue := make(map[string]float64)

	for _, booking := range bookings {
		flight := booking[1]
		price, _ := strconv.ParseFloat(booking[3], 64)

		flightRevenue[flight] += price
	}

	return flightRevenue
}

func CountSeatClassBookings(bookings [][]string) map[string]int {
	seatClassCount := make(map[string]int)

	for _, booking := range bookings {
		seatClass := booking[2]
		seatClassCount[seatClass]++
	}

	return seatClassCount
}

func FindMostRegularPassenger(bookings [][]string) (string, int) {
	passengerCount := make(map[string]int)

	for _, booking := range bookings {
		passengerName := booking[0]
		passengerCount[passengerName]++
	}

	maxBookings := 0
	var passenger string

	for name, count := range passengerCount {
		if maxBookings < count {
			maxBookings = count
			passenger = name
		}
	}

	return passenger, maxBookings
}

func FindHighestRevenueFlight(flightRevenue map[string]float64) (string, float64) {
	highestRevenue := 0.0
	var highestRevFlight string

	for flight, revenue := range flightRevenue {
		if highestRevenue < revenue {
			highestRevenue = revenue
			highestRevFlight = flight
		}
	}

	return highestRevFlight, highestRevenue
}
