package lru

import "fmt"

type LRUCache struct {
	data    map[int]int
	visited []int
}

func NewLRUCache() *LRUCache {
	return &LRUCache{
		data: map[int]int{
			0: 1, 1: 5, 2: 6, 3: 8, 4: 10, 5: 1,
		},
		visited: []int{},
	}
}

func (c *LRUCache) Add(key, value int) {
	if len(c.data) >= 6 {
		if len(c.visited) > 0 { // Prevent index out of range error
			leastUsed := c.visited[0]
			delete(c.data, leastUsed)
			c.visited = c.visited[1:]
			fmt.Println("Removed least recently used:", leastUsed)
		}
	}
	c.data[key] = value
	c.Visited(key)
	fmt.Println("Added Student ID:", key, "Score:", value)
}

func (c *LRUCache) Update(key, value int) {
	if _, exists := c.data[key]; exists {
		c.data[key] = value
		c.Visited(key)
		fmt.Println("Updated Student ID:", key, "New Score:", value)
	} else {
		fmt.Println("Student ID not found!")
	}
}

func (c *LRUCache) Visited(key int) {
	c.DeleteFromVisited(key)
	c.visited = append(c.visited, key)
}

func (c *LRUCache) Delete(key int) {
	if _, exists := c.data[key]; exists {
		delete(c.data, key)
		c.DeleteFromVisited(key)
		fmt.Println("Deleted Student ID:", key)
	} else {
		fmt.Println("Student ID not found!")
	}
}

func (c *LRUCache) DeleteFromVisited(key int) {
	for i, v := range c.visited {
		if v == key {
			c.visited = append(c.visited[:i], c.visited[i+1:]...)
			return
		}
	}
}
