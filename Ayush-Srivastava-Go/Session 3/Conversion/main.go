package main

import (
	"fmt"
	"strconv"
)

func main() {

}

// 1.
func stringSum(s1 string, s2 string, s3 string, s4 string, s5 string) int {

	n1, e1 := strconv.Atoi(s1)
	n2, e2 := strconv.Atoi(s2)
	n3, e3 := strconv.Atoi(s3)
	n4, e4 := strconv.Atoi(s4)
	n5, e5 := strconv.Atoi(s5)

	if(e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil){
		fmt.Println("All inputs must be intergers")
	}

	return (n1 + n2 + n3 + n4 + n5)

}

// 2.
func farToCelsius(temp int) float64 {

	cel := (temp - 32) * (5/9)

	return float64(cel)
}

// 3. 
func intToFloat(n1 int, n2 int, n3 int) float64 {

	f1 := float64(n1)
	f2 := float64(n2)
	f3 := float64(n3)

	avg := (f1 + f2 + f3)/3

	return avg
}


// 4.
func boolToInt(b bool) int {

	if b{
		return 1
	}else{
		return 0
	}
}

// 5.
func evenOdd(f float64) string{

	n := int(f)

	if(n % 2 == 0){
		return "even"
	}else{
		return "odd"
	}
}

// 6. 
func hexaToDec(h string) int {

	dec, err := strconv.ParseInt(h, 16, 64)

	if(err != nil){
		fmt.Println("error", err)
	}

	return int(dec)
}

// 7.
func sliceSum(arr []string) int {

	sum := 0

	for i:=0; i<len(arr); i++{
		n, err := strconv.Atoi(arr[i])

		if(err != nil){
			fmt.Println("Value ain't an integer!")
			break
		}
		
		sum += n
	}

	return sum
}

// 8.
func floatToSquare(s string) (float64, error) {

	f, err := strconv.ParseFloat(s, 64)

	if(err != nil){
		return 0, err
	}

	square := f * f

	return square, nil
}

// 9.
func Q9(num int, f float64, s string){

	convF := float64(num)
	convI := int(f)

	fmt.Printf("The converted float is %f and int is %d and string input was %s", convF, convI, s)
}

// 10.
func charToASCII(char string) (int, string) {

	if(len(char)>0){
		r := rune(char[0])
		
		i := int(r)

		return i, ""
	}

	return 0, "Input valid character!"

}