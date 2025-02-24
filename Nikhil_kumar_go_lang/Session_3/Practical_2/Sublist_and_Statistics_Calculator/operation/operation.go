package operation

func CalculateSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func CalculateAverage(numbers []int) float64 {
	return float64(CalculateSum(numbers)) / float64(len(numbers))
}

func FindMaximum(numbers []int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func FindMinimum(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}
