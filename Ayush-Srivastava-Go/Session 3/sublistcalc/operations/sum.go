package operations

import "fmt"

func Sum(arr []int) {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("Sum: ", sum)

}