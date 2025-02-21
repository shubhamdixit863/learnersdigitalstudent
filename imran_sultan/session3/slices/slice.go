package slices

import "fmt"

func SliceOpperation() {
	h := []int{9, 7}
	fmt.Println(h[0])
	fmt.Println(&h[0])
	fmt.Println(len(h), cap(h))

	h = append(h, 99)
	fmt.Println(&h[0])
	fmt.Println(len(h), cap(h))
	// h[3] = 10 cannot do that by direct but can use methods e.g (append ets)
}
