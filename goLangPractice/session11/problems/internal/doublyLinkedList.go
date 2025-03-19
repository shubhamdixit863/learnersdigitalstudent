package services

type DoublyLinkedList struct {
	Head  *Node
	Tail  *Node
	Size  int
	Limit int
}

func NewDoublyLinkedList(limit int) *DoublyLinkedList {
	return &DoublyLinkedList{
		Limit: limit,
	}
}

func (dll *DoublyLinkedList) Add(content string) {
	newNode := &Node{Content: content}

	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		dll.Tail.Next = newNode
		newNode.Prev = dll.Tail
		dll.Tail = newNode
	}

	dll.Size++

	if dll.Size > dll.Limit {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
		dll.Size--
	}
}
