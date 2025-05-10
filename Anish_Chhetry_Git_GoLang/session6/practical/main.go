package main

import (
	"fmt"
	"practical/createcache"
	"practical/get"
	"practical/put"
)

func main() {
	cacheSize := 2
	cache := createcache.NewLRUCache(cacheSize)

	cache = put.Put(1, 1, cache)
	fmt.Println("After Put(1, 1):", cache)

	cache = put.Put(2, 2, cache)
	fmt.Println("After Put(2, 2):", cache)

	value, cache := get.Get(1, cache)
	fmt.Println("Get(1):", value)
	fmt.Println("Cache after Get(1):", cache)

	cache = put.Put(3, 3, cache)
	fmt.Println("After Put(3, 3):", cache)

	value, cache = get.Get(2, cache)
	fmt.Println("Get(2):", value)
	fmt.Println("Cache after Get(2):", cache)

	cache = put.Put(4, 4, cache)
	fmt.Println("After Put(4, 4):", cache)

	value, cache = get.Get(1, cache)
	fmt.Println("Get(1):", value)
	fmt.Println("Cache after Get(1):", cache)

	value, cache = get.Get(3, cache)
	fmt.Println("Get(3):", value)
	fmt.Println("Cache after Get(3):", cache)

	value, cache = get.Get(4, cache)
	fmt.Println("Get(4):", value)
	fmt.Println("Cache after Get(4):", cache)
}
