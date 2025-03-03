package data

import "fmt"

var Capacity int
var Mapp = make(map[int]int, Capacity)
var Queue []int

func Get(key int) int {
	_, exist := Mapp[key]
	if !exist {
		fmt.Println("key not found: ", key)
		return -1
	} else {
		for i := 0; i < len(Queue); i++ {
			if Queue[i] == key {

				Queue = append(Queue[:i], Queue[i+1:]...)

			}
		}
		Queue = append(Queue, key)
		fmt.Println("[key ,value]",key, Mapp[key])
		// fmt.Println(Queue)
		

		return Mapp[key]

	}
}

func Put(key int, value int) {
	// fmt.Println("map", Mapp)
	_, exist := Mapp[key]

	if exist {
		Mapp[key] = value
		for i := 0; i < len(Queue); i++ {
			if Queue[i] == key {
				Queue = append(Queue[:i], Queue[i+1:]...)
				// fmt.Println(this.queue)
			}
		}
		Queue = append(Queue, key)
	} else {

		if len(Mapp) == Capacity {
			val := Queue[0]
			delete(Mapp, val)
			Mapp[key] = value
			Queue = append(Queue[1:], key)
		} else {
			Mapp[key] = value
			Queue = append(Queue, key)

		}

	}
	fmt.Println("current cache", Mapp)

}
