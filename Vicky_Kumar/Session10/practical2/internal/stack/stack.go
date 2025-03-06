package stack

import (
	"fmt"
)

type Stack struct {
	arr []int
	top int
}

func NewStack() *Stack {
	return &Stack{top: -1}
}

func (s *Stack) Push(num int) {
	s.arr = append(s.arr, num)
	s.top++
}
func (s *Stack) Pop() {
	if s.top == -1 {
		fmt.Println("Stack is empty")
		return
	}
	s.arr = s.arr[:s.top]
	s.top--
}
func (s *Stack) Top() int {
	if s.top == -1 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.arr[s.top]
}
func (s *Stack) Empty() bool {
	return s.top == -1
}
