package main

import (
	"fmt"
	"lru/operations"
)

func main() {
	cache := operations.Constructor(2)

	cache.Put(1, 10)
	cache.Put(2, 20)
	fmt.Println("Get 1:", cache.Get(1)) 

	cache.Put(3, 30)
	fmt.Println("Get 2:", cache.Get(2)) 

	fmt.Println("Get 3:", cache.Get(3)) 

	cache.Put(4, 40)
	fmt.Println("Get 1:", cache.Get(1)) 
	fmt.Println("Get 2:", cache.Get(2)) 
	fmt.Println("Get 4:", cache.Get(4)) 
}
