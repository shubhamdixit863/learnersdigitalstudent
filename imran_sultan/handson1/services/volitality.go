package services

import (
	"fmt"
	"math"
	"sort"
)

func MostVolatile(mp map[string][]float64) {
	mini := -math.MaxFloat64
	var ans string
	for key, value := range mp {
		sort.Float64s(value)
		num := value[len(value)-1] - value[0]
		if num > mini {
			mini = num
			ans = key
		}

	}

	fmt.Println("Most volatile stock is ", ans)

}
