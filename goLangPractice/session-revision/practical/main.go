package main

import (
	"fmt"
	"session-revision/practical/get"
	"session-revision/practical/set"
)

func main() {
	slc := make([]int, 5)
	slc[0] = 10
	slc[1] = 20
	slc[2] = 30
	slc[3] = 40
	slc[4] = 50

	fmt.Println(get.Get(slc, 20))
	fmt.Println(get.Get(slc, 40))
	fmt.Println(get.Get(slc, 70))
	fmt.Println(get.Get(slc, 20))
	fmt.Println(get.Get(slc, 30))
	fmt.Println(get.Get(slc, 50))

	set.Set(slc, 90)
	fmt.Println(slc)

	set.Set(slc, 110)
	fmt.Println(slc)
}
