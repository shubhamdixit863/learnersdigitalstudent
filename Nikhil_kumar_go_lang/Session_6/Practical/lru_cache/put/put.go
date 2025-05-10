// package put
package put

import (
	"lur_cache/priority"
)

func Put(key int, value string, LRU_Map map[int]string, Priority []int) []int {
	_, exists := LRU_Map[key]
	if exists {
		LRU_Map[key] = value
		Priority = priority.Change_priority(key, Priority)
		return Priority
	}

	if len(LRU_Map) < 3 {
		LRU_Map[key] = value
		Priority = priority.Change_priority(key, Priority)
	} else {
		if len(Priority) > 0 {
			delete(LRU_Map, Priority[0])
			Priority = priority.Remove_priority(Priority)
			LRU_Map[key] = value
			Priority = priority.Change_priority(key, Priority)
		}
	}
	return Priority
}

