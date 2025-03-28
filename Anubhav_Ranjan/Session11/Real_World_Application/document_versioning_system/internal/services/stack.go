package services

import "errors"

type Stack struct {
	nodes []*Node
}

func (s *Stack) Push(node *Node) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() (*Node, error) {
	if len(s.nodes) == 0 {
		return nil, errors.New("stack is empty")
	}
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return node, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.nodes) == 0
}

func (s *Stack) Clear() {
	s.nodes = []*Node{}
}
