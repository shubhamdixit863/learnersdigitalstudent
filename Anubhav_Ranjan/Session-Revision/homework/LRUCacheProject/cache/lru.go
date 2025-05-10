package cache

import "fmt"

var Cache = make(map[string]string)
var LRUtracker = []string{}
var Capacity int
var NextKey = 1
var FreedKeys []string

func InitLRUCache(size int) {
	if size <= 0 {
		fmt.Println("Invalid cache size, setting default capacity = 5")
		size = 5
	}
	Capacity = size

	// Reset Cache and Tracking Variables
	Cache = make(map[string]string)
	LRUtracker = []string{}
	NextKey = 1
	FreedKeys = []string{}
}
