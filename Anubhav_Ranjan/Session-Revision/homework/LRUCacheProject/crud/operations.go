package crud

import (
	"LRUCacheProject/cache"
	"LRUCacheProject/utils"
	"fmt"
	"strconv"
)

func Get(key string) (string, bool) {
	val, exists := cache.Cache[key]
	if !exists {
		return "", false
	}

	cache.LRUtracker = utils.UpdateLRUtracker(cache.LRUtracker, key)
	return val, true
}

func Put(key, value string) {

	if _, exists := cache.Cache[key]; exists {
		cache.Cache[key] = value
		cache.LRUtracker = utils.UpdateLRUtracker(cache.LRUtracker, key)
		return
	}

	if len(cache.LRUtracker) >= cache.Capacity {
		utils.RemoveLRU()
	}

	cache.Cache[key] = value
	cache.LRUtracker = append(cache.LRUtracker, key)
}

func Post(value string) string {
	var newKey string

	if len(cache.FreedKeys) > 0 {
		newKey = cache.FreedKeys[0]
		cache.FreedKeys = cache.FreedKeys[1:]
	} else {
		newKey = strconv.Itoa(cache.NextKey)
		cache.NextKey++
	}

	if len(cache.LRUtracker) >= cache.Capacity {
		utils.RemoveLRU()
	}

	cache.Cache[newKey] = value
	cache.LRUtracker = append(cache.LRUtracker, newKey)

	fmt.Println("Added:", newKey, "=>", value)
	return newKey
}

func Delete(key string) {
	if _, exists := cache.Cache[key]; !exists {
		fmt.Println("Key not found in cache")
		return
	}

	delete(cache.Cache, key)
	cache.LRUtracker = utils.RemoveKeyFromSlice(cache.LRUtracker, key)

	cache.FreedKeys = append(cache.FreedKeys, key)

}
