package services

func Aggregator(mp *map[string]int, r map[string]int) {
	finalCounts := *mp

	for level, count := range r {
		finalCounts[level] += count
	}
}
