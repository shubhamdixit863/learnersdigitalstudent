package createcache

type LRUCache struct {
	Size  int
	Cache map[int]int
	Key   []int
}

func NewLRUCache(size int) LRUCache {
	return LRUCache{
		Size:  size,
		Cache: make(map[int]int),
		Key:   make([]int, 0, size),
	}
}
