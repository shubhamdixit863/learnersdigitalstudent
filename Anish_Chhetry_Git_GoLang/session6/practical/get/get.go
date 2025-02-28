package get

import (
	"practical/createcache"
	"practical/update"
)

func Get(key int, cache createcache.LRUCache) (int, createcache.LRUCache) {
	for k := range cache.Cache {
		if k == key {
			cache = update.Update(k, cache)
			return k, cache
		}
	}
	return -1, cache
}
