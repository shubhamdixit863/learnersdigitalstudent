package update

import (
	"practical/createcache"
)

func Update(key int, cache createcache.LRUCache) createcache.LRUCache {
	NewKey := make([]int, 0, len(cache.Key))
	for _, x := range cache.Key {
		if x != key {
			NewKey = append(NewKey, x)
		}
	}
	NewKey = append(NewKey, key)
	cache.Key = NewKey
	return cache
}
