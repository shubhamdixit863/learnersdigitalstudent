package services

func AggregateLogs(processedChannel <-chan LogEntry, countChannel chan<- map[string]int) {
	counts := make(map[string]int)
	for entry := range processedChannel {
		counts[entry.Level]++
	}
	countChannel <- counts
}
