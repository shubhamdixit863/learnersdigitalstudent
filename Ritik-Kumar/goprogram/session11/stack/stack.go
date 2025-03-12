package stack

import "session11/models"

type Stack struct {
	items []*models.Node
}

func (s *Stack) Push(node *models.Node) {
	s.items = append(s.items, node)
}

func (s *Stack) Pop() *models.Node {
	if len(s.items) == 0 {
		return nil
	}
	lastIndex := len(s.items) - 1
	node := s.items[lastIndex]
	s.items = s.items[:lastIndex] 
	return node
}

func (s *Stack) Clear() {
	s.items = nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
