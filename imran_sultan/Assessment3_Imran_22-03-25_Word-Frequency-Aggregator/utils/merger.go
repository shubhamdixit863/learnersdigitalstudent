package utils

func MergeResults(in <-chan map[string]int) map[string]int {
	finalFreq := make(map[string]int)

	for freq := range in {
		for word, count := range freq {
			finalFreq[word] += count
		}
	}

	return finalFreq
}
