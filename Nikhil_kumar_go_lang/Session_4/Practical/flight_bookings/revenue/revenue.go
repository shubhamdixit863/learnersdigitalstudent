package revenue

import (
	"fmt"
	"strconv"
	"strings"
)

func Revenue(book []string) {
	revenuePerFlight := make(map[string]float64)

	for _, bookin := range book {
		comp := strings.Split(bookin, ":")
		amount, err := strconv.ParseFloat(comp[3], 64)
		if err != nil {
			continue
		}
		revenuePerFlight[comp[1]] += amount
	}
	for k, v := range revenuePerFlight {
		fmt.Printf("flight : %s revenue : %.2f\n", k, v)
	}
	// fmt.Println("revenue", revenuePerFlight)

}
func Highest_revenue(book []string) {
	revenuePerFlight := make(map[string]float64)
	var max float64
	var flight string
	for _, bookin := range book {
		comp := strings.Split(bookin, ":")
		amount, err := strconv.ParseFloat(comp[3], 64)
		if err != nil {
			continue
		}
		revenuePerFlight[comp[1]] += amount
	}
	for k, v := range revenuePerFlight {
		if v > max {
			max = v
			flight = k
		}
	}
	fmt.Printf("Flight with highest revenue: %s with revenue: %.2f\n", flight, max)

}
