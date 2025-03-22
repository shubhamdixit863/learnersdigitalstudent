package storage

import "sync"

// Storage maps words to the links where they appear
type Storage struct {
	mu   sync.Mutex
	data map[string]map[string]int
}

// NewStorage initializes storage
func NewStorage() *Storage {
	return &Storage{data: make(map[string]map[string]int)}
}

// Add increments the count of a word found on a link
func (s *Storage) Add(word, link string, count int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[word]; !exists {
		s.data[word] = make(map[string]int)
	}
	s.data[word][link] += count
}

// GetData returns stored data
func (s *Storage) GetData() map[string]map[string]int {
	return s.data
}
