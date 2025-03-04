package utils

import (
	"LRUCacheProject/cache"
)

func UpdateLRUtracker(slice []string, key string) []string {
	slice = RemoveKeyFromSlice(slice, key)
	return append(slice, key)
}

func RemoveKeyFromSlice(slice []string, key string) []string {
	for i, k := range slice {
		if k == key {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func RemoveLRU() {
	if len(cache.LRUtracker) == 0 {
		return
	}

	lruKey := cache.LRUtracker[0]
	cache.LRUtracker = cache.LRUtracker[1:]

	delete(cache.Cache, lruKey)

	cache.FreedKeys = append(cache.FreedKeys, lruKey)
}
