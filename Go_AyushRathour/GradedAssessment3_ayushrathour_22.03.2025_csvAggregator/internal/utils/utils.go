package utils

import (
	"assessment3/internal/types"
	"sync"
)

func MergeMaps(dest, src types.UserSpend, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	for user, spend := range src {
		dest[user] += spend
	}
}
