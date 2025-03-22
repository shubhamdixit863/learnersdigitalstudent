package services

import (
	"fmt"
	"sehandson1/data"
	"sort"
)

func MInMax(mp map[string][]float64) {

	for key, value := range mp {
		sort.Float64s(value)
		data.Min[key] = value[0]
		data.Max[key] = value[len(value)-1]
	}
	fmt.Println("Minimum values/stock: ")
	for key, value := range data.Min {
		fmt.Println(key, value)

	}

	fmt.Println("Maximum values/stock: ")
	for key, value := range data.Max {
		fmt.Println(key, value)

	}

}
