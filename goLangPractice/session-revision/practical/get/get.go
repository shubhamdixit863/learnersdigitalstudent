package get

import "session-revision/practical/lru"

func Get(slc []int, key int) int {
	for i, num := range slc {
		if num == key {
			lru.AddIdx(i)
			return i
		}
	}
	return -1
}
