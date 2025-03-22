package services

import (
	"fmt"
	"sehandson1/data"
)

func CalculateAvg(mp map[string][]float64) {
	for key, value := range mp {
		var sum float64
		sum = 0
		for i := 0; i < len(value); i++ {
			sum = sum + value[i]
		}
		avg := (float64(sum)) / (float64(len(value)))
		data.Average[key] = avg
	}
	fmt.Println("Average values/stock: ")
	for key, value := range data.Average {
		fmt.Println(key, value)

	}

}
