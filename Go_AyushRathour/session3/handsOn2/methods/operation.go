package methods

import("fmt")

func Operations(subList []int){
	fmt.Printf("SubList: %v\n", subList)

	sum:=0;
	max:=subList[0];
	for i:=0; i<len(subList); i++{
		sum+=subList[i]
		if subList[i]>max{
			max=subList[i]
		}
	}
	fmt.Printf("Sum: %d\n", sum)
    average:=float64(sum)/float64(len(subList))
    fmt.Printf("Average: %f\n", average)
	fmt.Printf("Max: %d\n", max)

}