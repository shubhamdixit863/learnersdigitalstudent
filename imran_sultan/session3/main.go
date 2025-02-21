package main

import (
	"fmt"
	"math"
	"reflect"
	"session3/pointers"
	"session3/slices"
	"strconv"
)

func main() {
	// c := 9
	// d := 9
	// f := float32(d)
	// var g float32
	// g = 9.78
	// i := int(g)
	// fmt.Println(reflect.TypeOf(i))
	// fmt.Println(i)
	// fmt.Println(g)
	// fmt.Println(reflect.TypeOf(g))
	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(d))
	// fmt.Println(reflect.TypeOf(f))
	// fmt.Println(("hello"))
	// var b bool
	// fmt.Println(b)
	// in := int(b)
	// h, err := strconv.Atoi("A")
	// if err != nil {
	// 	fmt.Println("conversion not possible", err)

	// }
	// fmt.Println(h)
	// h1, err := strconv.Atoi("2")
	// if err != nil {
	// 	fmt.Println("conversion not possible", err)

	// }
	// fmt.Println(h1)
	// fmt.Println(reflect.TypeOf(h1))
	// var str = "imran"
	// a1 := string('2')
	// fmt.Println(a1)
	// j := strconv.Itoa(2)
	// fmt.Println(reflect.TypeOf(j))
	// strToint("123", "456", "786", "890", "656")
	// intTofloat(12)
	// boolToint(true)
	// floatToint(11.2)
	// hexaTodecimal("1A3F")
	// HexatoD("1A")
	// str := []string{"2", "2"}
	// strToin(str)
	// str := "112.34"
	// h, _ := strconv.ParseFloat(str, 32)
	// fmt.Println(h)
	pointers.Hello()
	slices.SliceOpperation()

}
func strToint(s1 string, s2 string, s3 string, s4 string, s5 string) {
	num1, err1 := strconv.Atoi(s1)
	num2, err2 := strconv.Atoi(s2)
	num3, err3 := strconv.Atoi(s3)
	num4, _ := strconv.Atoi(s4)
	num5, _ := strconv.Atoi(s5)
	fmt.Println(num1, err1)
	fmt.Println(num2, err2)
	fmt.Println(num3, err3)
	fmt.Println(num4)
	fmt.Println(num5)

}
func intTofloat(num int) {
	f := float32(num)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
}
func boolToint(b bool) {
	var n int
	if b {
		n = 1
	} else {
		n = 0
	}
	fmt.Println(n)

}
func floatToint(f float32) {
	num := int(f)
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("odd")
	}

}
func hexaTodecimal(str string) {
	num, _ := strconv.ParseInt(str, 16, 32)
	fmt.Println(num)

}
func HexatoD(str string) {
	var place float64 = 0
	var ans float64 = 0
	for i := len(str) - 1; i >= 0; i-- {
		x := str[i]
		h1, err := strconv.Atoi(string(x))
		if err != nil {
			if x == 'A' {
				fmt.Println("i am here ")
				ans += 10 * (math.Pow(16, float64(place)))
				fmt.Println(ans)
				place++
			}
			if x == 'B' {
				ans += 11 * (math.Pow(16, float64(place)))
				place++
			}
			if x == 'C' {
				ans += 12 * (math.Pow(16, float64(place)))
				place++
			}
			if x == 'D' {
				ans += 13 * (math.Pow(16, float64(place)))
				place++
			}
			if x == 'E' {
				ans += 14 * (math.Pow(16, float64(place)))
				place++
			}
			if x == 'F' {
				ans += 15 * (math.Pow(16, float64(place)))
				place++
			}

		} else {
			fmt.Println("i ah in else")
			ans += float64(h1) * (math.Pow(16, float64(place)))
			fmt.Print(ans)
			place++
		}
	}
	fmt.Println(ans)
}
func strToin(str []string) {
	var sum = 0
	for i := 0; i < len(str); i++ {
		h1, _ := strconv.Atoi(str[i])
		sum += h1
	}
	fmt.Println(sum)

}
