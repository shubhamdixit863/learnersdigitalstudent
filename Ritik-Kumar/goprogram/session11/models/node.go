package models

type Node struct {
	Content string
	Prev    *Node
	Next    *Node
}
