package services

func StoreLogs(processedChannel chan LogEntry, resultLog chan map[string]int){
	logCounts:=make(map[string]int)
	for log := range processedChannel {
		logCounts[log.Level]++
	}

	resultLog <- logCounts
	close(resultLog)
}
