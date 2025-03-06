package doublylinkedlist

import (
	"fmt"
)

type DoublyLinkedList struct {
	Head *Node
	Size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Size: 0,
	}
}

type Node struct {
	Data string
	Prev *Node
	Next *Node
}

func NewNode(s string) *Node {
	return &Node{Data: s, Prev: nil, Next: nil}
}

func (l *DoublyLinkedList) Insert(str string, pos int) {
	newNode := NewNode(str)
	if pos == 0 {
		l.Head = newNode
	} else {
		curr := l.Head
		for i := 0; i < pos-1; i++ {
			curr = curr.Next
		}
		curr.Next = newNode
		newNode.Prev = curr
	}
	l.Size++
}
func (l *DoublyLinkedList) GetData(pos int) string {
	if pos < 0 || pos >= l.Size {
		fmt.Println("position out of bound")
		return ""
	}
	currNode := l.Head
	for i := 0; i < pos; i++ {
		currNode = currNode.Next
	}
	return currNode.Data
}
func (l *DoublyLinkedList) Delete(pos int) {
	if pos < 0 || pos >= l.Size {
		fmt.Println("position out of bound")
		return
	}
	if pos == 0 {
		l.Head = l.Head.Next
		l.Head.Prev = nil
	} else {
		currNode := l.Head
		for i := 0; i < pos; i++ {
			currNode = currNode.Next
		}
		currNode.Prev = currNode.Next
		currNode.Next.Next = currNode.Prev
	}
	l.Size--
}
func (l *DoublyLinkedList) PrintList() {
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Data, " <-> ")
		curr = curr.Next
	}
	fmt.Println("nil")
}
