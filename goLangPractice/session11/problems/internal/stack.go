package services

type Stack struct {
	Items []*Node
}

func (s *Stack) Push(node *Node) {
	s.Items = append(s.Items, node)
}

func (s *Stack) Pop() *Node {
	if len(s.Items) == 0 {
		return nil
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}

func (s *Stack) Clear() {
	s.Items = nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}
