// package display
package display

import (
	"fmt"
	"lur_cache/get"
	"lur_cache/priority"
	"lur_cache/put"
)

const Capacity int = 3

var LRU_Map = make(map[int]string)
var Priority = make([]int, 0)

func Display() {
	var key int
	var value string
	for {
		var option int
		fmt.Println("1 for put \n2 for get \n3 for exit \n4 for check LRU \n5 for Priority")
		fmt.Scan(&option)
		if option == 1 {
			fmt.Println("enter key and value")
			fmt.Println("enter key ")
			fmt.Scan(&key)
			fmt.Println("enter value")
			fmt.Scan(&value)
			Priority = put.Put(key, value, LRU_Map, Priority)
		} else if option == 2 {
			fmt.Println("enter key")
			fmt.Scan(&key)
			Priority = get.Get(LRU_Map, key, Priority)
		} else if option == 3 {
			break
		} else if option == 4 {
			show_map(LRU_Map)
		} else if option == 5 {
			priority.Show_priority(Priority)
		} else {
			fmt.Println("wrong input")
		}
	}
}

func show_map(LRU_Map map[int]string) {
	fmt.Println(LRU_Map)
}
