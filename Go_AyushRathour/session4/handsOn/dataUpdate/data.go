package dataUpdate

import (
	"strconv"
	"strings"
)

func AddData(bookings []string) (map[string]int, map[string]float64, map[string]int) {
	name_map := make(map[string]int)
	flightName_map := make(map[string]float64)
	seatclass_map := make(map[string]int)

	for i := 0; i < len(bookings); i++ {
		data := strings.Split(bookings[i], ":")
		name := data[0]
		flightName := data[1]
		seatclass := data[2]
		price, _ := strconv.ParseFloat(data[3], 64)

		if name_map[name] == 0 {
			name_map[name] = 1
		} else {
			name_map[name]++
		}

		if flightName_map[flightName] == 0 {
			flightName_map[flightName] = price
		} else {
			flightName_map[flightName] += price
		}

		if seatclass_map[seatclass] == 0 {
			seatclass_map[seatclass] = 1
		} else {
			seatclass_map[seatclass]++
		}
	}
	return name_map, flightName_map, seatclass_map

}
