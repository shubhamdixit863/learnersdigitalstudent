package put

import (
	"practical/createcache"
	"practical/update"
)

func Put(key int, value int, cache createcache.LRUCache) createcache.LRUCache {
	for x := range cache.Cache {
		if x == key {
			cache.Cache[key] = value
			update.Update(key, cache)
			return cache
		}
	}
	if cache.Size == len(cache.Cache) {
		delete(cache.Cache, cache.Key[0])
		cache.Key = cache.Key[1:]
	}
	cache.Cache[key] = value
	cache.Key = append(cache.Key, key)
	return cache
}
