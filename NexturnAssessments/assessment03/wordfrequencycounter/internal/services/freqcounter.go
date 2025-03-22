package services


func CalculateFrequency(words []string) map[string]int {
	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}
	return freqMap
}
