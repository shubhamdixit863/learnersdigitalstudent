package lru

import "fmt"

var temp []int

func AddIdx(idx int) {
	for i, num := range temp {
		if idx == num {
			temp = append(temp[:i], temp[i+1:]...)
		}
	}
	temp = append(temp, idx)
}

func Lru(slc []int, value int) {
	if len(temp) == 0 {
		fmt.Println("There is no Least Recently used element")
	} else {
		num := temp[0]
		temp = append(temp[1:])
		temp = append(temp, num)
		slc[num] = value
	}
}
