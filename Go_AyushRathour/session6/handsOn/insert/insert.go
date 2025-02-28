package insert

import (
	"fmt"
)

var LRU []int
var m = make(map[int]int)
var Capacity int

func Put(key int, value int) {
	fmt.Println(LRU)
	_, exist := m[key]

	if exist {
		m[key] = value
		UpdateLRU(key)

	} else {

		if len(m) == 2 {
			val := LRU[0]
			delete(m, val)
			m[key] = value
			LRU = LRU[1:]
			fmt.Println(LRU)
			UpdateLRU(key)

		} else {
			m[key] = value
			UpdateLRU(key)
		}

	}
	fmt.Println(LRU)

}
func Get(key int) int {
	for k, v := range m {
		if k == key {
			UpdateLRU(key)
			return v
		}
	}
	return -1
}

func UpdateLRU(key int) {
	index := -1
	for i, v := range LRU {
		if v == key {
			index = i
			break
		}
	}
	if index != -1 {
		LRU = append(LRU[:index], LRU[index+1:]...)
	}
	LRU = append(LRU, key)



}
