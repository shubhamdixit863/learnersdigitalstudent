package operations

import "fmt"

func Average(arr []int) {
	var avg float64
	sum := 0
	for _, num := range arr {
		sum += num
	}
	avg = float64(sum) / float64(len(arr))
	fmt.Println("Average: ", avg)

}
