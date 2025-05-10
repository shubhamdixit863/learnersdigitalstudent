package bookings

import (
	"fmt"
	"strconv"
	"strings"
)

var Bookings = []string{
	"Alice:FL123:Economy:120.50",
	"Bob:FL123:Business:450.00",
	"Charlie:FL456:Economy:150.75",
	"Alice:FL456:Economy:150.75",
	"Bob:FL123:Economy:120.50",
}
var details = [][]string{}

const idxName int = 0
const idxFlight int = 1
const idxSeat int = 2
const idxPrice int = 3

func Split() {
	for i := 0; i < len(Bookings); i++ {
		detail := strings.Split(Bookings[i], ":")
		if len(detail) != 4 {
			continue
		}
		details = append(details, detail)
	}
	fmt.Println(details)
}

func TotalRevenue() map[string]float64 {
	m := make(map[string]float64)
	for i := 0; i < len(details); i++ {
		revenue, err := strconv.ParseFloat(details[i][idxPrice], 64)
		if err != nil {
			fmt.Println(err)
			return m
		}
		m[details[i][idxFlight]] += revenue
	}
	return m
}
func SeatCount() map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(details); i++ {
		m[details[i][idxFlight]]++
	}
	return m
}
func MostBooking() (string, int) {
	m := make(map[string]int)
	for i := 0; i < len(details); i++ {
		m[details[i][idxName]]++
	}
	name := ""
	maxBooking := 0
	for k, v := range m {
		if v > maxBooking {
			name = k
			maxBooking = v
		}
	}
	return name, maxBooking
}
func MostRevenueFlight() (string, float64) {

	m := TotalRevenue()
	name := ""
	maxRevenue := 0.0
	for k, v := range m {
		if v > maxRevenue {
			name = k
			maxRevenue = v
		}
	}
	return name, maxRevenue
}
