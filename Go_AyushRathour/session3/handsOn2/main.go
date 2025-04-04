package main

import("fmt"
		"handsOn/slice"
	    "handsOn/methods"
	)


func main() {
	fmt.Printf("Welcome to the sublist Calculator \n")

	// Taking input from the user
	
	fmt.Printf("Enter the elements of the list : ")


	var num []int
	var n int
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		num = append(num, n)
	}
	fmt.Printf("\n")


	Operation(num)

}

func Operation(num []int) {
	

	fmt.Printf("Enter the start index: ")
	var start int
	fmt.Scanln(&start)
	fmt.Printf("Enter the end index : ")
	var end int
	fmt.Scanln(&end)
	fmt.Printf("\n")
	sublist := slice.Slice(num, start, end)
	if sublist == nil {
		return
	}


	methods.Operations(sublist)

	fmt.Printf("Do you want to slice again? (yes/no): ")
	var choice string
	fmt.Scanln(&choice)
	if choice == "yes" {
		Operation(num)
	} else {
		fmt.Println("Exiting...")
		return
	}

}