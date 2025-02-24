package display

import (
	bookingbyseat "flight_booking/booking_by_seat"
	"flight_booking/revenue"
	"fmt"
)

func Display(book []string) {
	fmt.Println("Revenue per flight:")
	revenue.Revenue(book)
	fmt.Println("Bookings by seat class:")
	bookingbyseat.Booking_by_seat(book)
	fmt.Println("Passenger with the most bookings:")
	bookingbyseat.Passenger_with_most_booking(book)
	fmt.Println("Flight with the highest revenue:")
	revenue.Highest_revenue(book)
}
