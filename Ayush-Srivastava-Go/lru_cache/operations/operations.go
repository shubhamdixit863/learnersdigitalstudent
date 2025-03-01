package operations

type LRUCache struct {
	capacity int
	cache    map[int]int
	vis      []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
		vis:      []int{},
	}
}

func (cache *LRUCache) moveToEnd(key int) {
	for idx, k := range cache.vis {
		if k == key {
			cache.vis = append(cache.vis[:idx], cache.vis[idx+1:]...)
			break
		}
	}
	cache.vis = append(cache.vis, key)
}

func (cache *LRUCache) Get(key int) int {
	if val, ok := cache.cache[key]; ok {
		cache.moveToEnd(key)
		return val
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if _, ok := cache.cache[key]; ok {
		cache.moveToEnd(key)
	} else {
		if len(cache.cache) >= cache.capacity {
			deleteKey := cache.vis[0]
			delete(cache.cache, deleteKey)
			cache.vis = cache.vis[1:]
		}
		cache.vis = append(cache.vis, key)
	}
	cache.cache[key] = value
}
