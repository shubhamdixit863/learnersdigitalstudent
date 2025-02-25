package analyze

import (
	"fmt"
	"strconv"
	"strings"
)

func AnalyzeBookings(bookings []string) {
	flightRevenue := make(map[string]float64)
	seatClassCounts := make(map[string]int)
	passengerBookings := make(map[string]int)
	booked := make(map[string]bool)

	for _, b := range bookings {
		parts := strings.Split(b, ":")
		if len(parts) != 4 {
			continue
		}

		n, f, s, p := parts[0], parts[1], parts[2], parts[3]

		price, err := strconv.ParseFloat(p, 64)
		if err != nil || price < 0 {
			continue
		}

		key := fmt.Sprintf("%s:%s:%s", n, f, s)
		if booked[key] {
			continue
		}
		booked[key] = true

		flightRevenue[f] += price
		seatClassCounts[s]++
		passengerBookings[n]++
	}

	fmt.Println("Revenue per flight:")
	for f, r := range flightRevenue {
		fmt.Printf("%s: $%.2f\n", f, r)
	}

	fmt.Println("Bookings by seat class:")
	for s, c := range seatClassCounts {
		fmt.Printf("%s: %d\n", s, c)
	}

	mf := ""
	mb := 0
	for n, c := range passengerBookings {
		if c > mb {
			mf = n
			mb = c
		}
	}

	fmt.Printf("Passenger with the most bookings: %s (%d bookings)\n", mf, mb)

	hrf := ""
	mr := 0.0
	for f, r := range flightRevenue {
		if r > mr {
			hrf = f
			mr = r
		}
	}

	fmt.Printf("Flight with the highest revenue: %s ($%.2f)\n", hrf, mr)
}
