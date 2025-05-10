package operations

import "fmt"

func Max(arr []int) {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println("Maximum: ", max)

}
