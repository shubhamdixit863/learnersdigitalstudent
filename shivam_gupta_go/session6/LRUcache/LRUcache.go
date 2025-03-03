package LRUcache

import "fmt"

type LRUcache1 struct {
	capacity int
	mp       map[string]int
	order    []string
}

var LRUcache = LRUcache1{
	capacity: 3,
	mp:       make(map[string]int),
	order:    []string{},
}

func MoveLast(key string) {
	for i, k := range LRUcache.order {
		if k == key {
			LRUcache.order = append(LRUcache.order[:i], LRUcache.order[i+1:]...)
			break
		}
	}
	LRUcache.order = append(LRUcache.order, key)
}

func Get(key string) int {
	if val, ok := LRUcache.mp[key]; ok {
		MoveLast(key)
		return val
	}
	return -1
}

func Put(key string, value int) {
	if _, ok := LRUcache.mp[key]; ok {
		LRUcache.mp[key] = value
		MoveLast(key)
		return
	}

	if len(LRUcache.order) >= LRUcache.capacity {
		temp := LRUcache.order[0]
		delete(LRUcache.mp, temp)
		LRUcache.order = LRUcache.order[1:]
	}

	LRUcache.mp[key] = value
	LRUcache.order = append(LRUcache.order, key)
}

func Printcache() {
	fmt.Println("cache : ", LRUcache.mp)
	fmt.Println("Order : ", LRUcache.order)
}

func LRUCache() {

	Put("A", 10)
	Put("B", 20)
	Printcache()
	fmt.Println(Get("A"))

	Printcache()
	Put("C", 30)
	Printcache()

	fmt.Println(Get("B"))
	Put("D", 40)
	Printcache()

	Put("C", 30)
	Printcache()
}


