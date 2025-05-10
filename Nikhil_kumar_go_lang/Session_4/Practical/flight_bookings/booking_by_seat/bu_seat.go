package bookingbyseat

import (
	"fmt"
	"strings"
)

func Booking_by_seat(book []string) {
	bookingsBySeatClass := make(map[string]int)

	for _, bookin := range book {
		comp := strings.Split(bookin, ":")
		bookingsBySeatClass[comp[2]]++
	}
	for k, v := range bookingsBySeatClass {
		fmt.Printf("class %s total %d \n", k, v)
	}
	// fmt.Println(bookingsBySeatClass)
}

func Passenger_with_most_booking(book []string) {
	passengerBookings := make(map[string]int)
	for _, bookin := range book {
		comp := strings.Split(bookin, ":")
		passengerBookings[comp[0]]++
	}
	for k, v := range passengerBookings {
		fmt.Printf("name : %s booking %d\n", k, v)
	}

}
