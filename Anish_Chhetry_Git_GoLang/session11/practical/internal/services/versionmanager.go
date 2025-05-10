package services

import (
	"fmt"
	"log"
)

type Node struct {
	prev *Node
	next *Node
	data string
}
type VersionManager struct {
	limit   int
	size    int
	current *Node
	head    *Node
	stack   []*Node
}

func (v *VersionManager) AddVersion(content string) {
	NewNode := &Node{data: content}
	if v.size == 0 {
		v.current = NewNode
		v.head = NewNode
		log.Println("Created a New Head Node")
	} else {
		v.current.next = NewNode
		NewNode.prev = v.current
		v.current = NewNode
		log.Println("Created a New Node")
	}
	v.size++
	if v.size > v.limit {
		v.head.next = v.head
		v.head.prev = nil
	}
}

func (v *VersionManager) GetCurrentVersion() string {
	return v.current.data
}

func (v *VersionManager) Undo() string {
	v.stack = append(v.stack, v.current)
	v.current = v.current.prev
	v.current.next = nil
	return fmt.Sprintf("Current Node after Undo: %s", v.current.data)
}

func (v *VersionManager) Redo() string {
	top := v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
	v.current.next = top
	top.prev = v.current
	v.current = top

	return fmt.Sprintf("Current Node after Redo: %s", v.current.data)

}

func NewVersionManager(limit int) *VersionManager {
	return &VersionManager{
		limit:   limit,
		size:    0,
		current: nil,
		head:    nil,
		stack:   make([]*Node, 0),
	}
}
