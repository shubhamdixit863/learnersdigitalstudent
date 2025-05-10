package database

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var Bookings = []string{
	"Alice:FL123:Economy:120.50",
	"Bob:FL123:Business:450.00",
	"Charlie:FL456:Economy:150.75",
	"Alice:FL456:Economy:150.75",
	"Bob:FL123:Economy:120.50",
	"Bob:FL123:Economy:120.50",
	"Invalid:FL789:First:-100.00",
	"John:FL123::150.00",
	"JackFL123Economy120.50",
}

var bookingRegex = regexp.MustCompile(`^[A-Za-z]+:[A-Z]{2}\d{3}:(Economy|Business|First):\d+(\.\d{1,3})?$`)

func GetBookings() [][]string {
	var validBookings [][]string
	uniqueBookings := make(map[string]bool)

	for _, booking := range Bookings {
		if !bookingRegex.MatchString(booking) {
			fmt.Println("Skipped invalid format:", booking)
			continue
		}

		booking_data := strings.Split(booking, ":")
		if len(booking_data) != 4 {
			fmt.Println("Skipped Missing Data Format:", booking)
			continue
		}

		name, flight, seatClass, priceStr := booking_data[0], booking_data[1], booking_data[2], booking_data[3]

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil || price < 0 {
			fmt.Println("Skipped Invalid Price:", booking)
			continue
		}

		bookingKey := name + ":" + flight + ":" + seatClass
		if uniqueBookings[bookingKey] {
			fmt.Println("Skipped Duplicate Booking:", booking)
			continue
		}

		validBookings = append(validBookings, booking_data)
		uniqueBookings[bookingKey] = true

	}

	return validBookings
}
