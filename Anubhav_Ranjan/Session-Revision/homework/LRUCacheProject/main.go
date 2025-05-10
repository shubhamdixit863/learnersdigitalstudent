package main

import (
	"LRUCacheProject/cache"
	"LRUCacheProject/crud"
	"fmt"
)

func main() {
	cache.InitLRUCache(5)

	crud.Post("A")
	fmt.Println("LRU Tracker:", cache.LRUtracker)
	crud.Post("B")
	fmt.Println("LRU Tracker:", cache.LRUtracker)
	crud.Post("C")
	fmt.Println("LRU Tracker:", cache.LRUtracker)
	fmt.Println("Cache:", cache.Cache)

	value, _ := crud.Get("2")
	fmt.Println("Get(2):", value)
	fmt.Println("LRU Tracker:", cache.LRUtracker)
	fmt.Println("Cache:", cache.Cache)

	crud.Put("1", "D")
	fmt.Println("LRU Tracker:", cache.LRUtracker)
	fmt.Println("Cache:", cache.Cache)

	crud.Post("New")
	fmt.Println("LRU Tracker:", cache.LRUtracker)
	fmt.Println("Cache:", cache.Cache)

	crud.Delete("3")
	fmt.Println("LRU Tracker:", cache.LRUtracker)
	fmt.Println("Cache:", cache.Cache)
}
