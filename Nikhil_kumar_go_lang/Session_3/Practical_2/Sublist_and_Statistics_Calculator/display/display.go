package display

import (
	"Sublist_and_Statistics_Calculator/operation"
	"Sublist_and_Statistics_Calculator/slicesNum"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Display() {
	fmt.Println("Enter a list of numbers (separated by spaces): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numbers := strings.Split(strings.TrimSpace(input), " ")
	var intSlice []int
	for _, num := range numbers {
		i, _ := strconv.Atoi(num)
		intSlice = append(intSlice, i)
	}
	fmt.Println(intSlice)

	var start int
	var end int
	for {
		var input string
		fmt.Println("Enter the start index: ")
		fmt.Scan(&start)
		fmt.Println("Enter the end index: ")
		fmt.Scan(&end)

		sl := slicesNum.SliceNumbers(intSlice, start, end)
		fmt.Println("Sliced Sublist:", sl)
		fmt.Println("Sum", operation.CalculateSum(sl))
		fmt.Println("Average", operation.CalculateAverage(sl))
		fmt.Println("Maximmum", operation.FindMaximum(sl))
		fmt.Println("Minimum", operation.FindMinimum(sl))

		fmt.Println("Do you want to slice again? (yes/no):")
		fmt.Scan(&input)
		if input == "no" {
			fmt.Println("Exiting the application. Goodbye!")
			break
		}
	}
}
