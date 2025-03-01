package set

import (
	"session-revision/practical/lru"
)

func Set(slc []int, value int) {
	if len(slc) >= 5 {
		lru.Lru(slc, value)
	} else {
		slc = append(slc, value)
	}
}
