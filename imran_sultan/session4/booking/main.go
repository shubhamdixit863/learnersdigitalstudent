package booking

import (
	"strings"
)

var Bookings = []string{
	"Alice:FL123:Economy:120.50",
	"Bob:FL123:Business:450.00",
	"Charlie:FL456:Economy:150.75",
	"Alice:FL456:Economy:150.75",
	"Bob:FL123:Economy:120.50",
}

var result = make([]string, len(Bookings[0]))

func Sepprating() {
	m1 := make(map[string]int)
	for i := 0; i < len(Bookings); i++ {
		result = strings.Split(Bookings[i], ":")
		m1[result[2]] = result[3]

	}

}
func Mapping(name []string, fligt []string, charlielass []string, revenue []int) {

}
