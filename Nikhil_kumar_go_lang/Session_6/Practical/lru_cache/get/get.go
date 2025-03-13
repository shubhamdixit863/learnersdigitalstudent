// package get
package get

import (
	"fmt"
	"lur_cache/priority"
)

func Get(LRU_Map map[int]string, key int, Priority []int) []int {
	value, exists := LRU_Map[key]
	if exists {
		fmt.Println(value)
		Priority = priority.Change_priority(key, Priority)
	} else {
		fmt.Println("not found")
	}
	return Priority
}
