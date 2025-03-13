package docversiontrack

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

type DocVersionTrack struct {
	Head *Node
	Tail *Node
}

func NewDocVersionTrack() *DocVersionTrack {
	return &DocVersionTrack{}
}

func (vt *DocVersionTrack) Append(item interface{}) *Node {
	newVersion := &Node{Value: item}
	if vt.Head == nil {
		vt.Head = newVersion
	}
	if vt.Tail != nil {
		vt.Tail.Right = newVersion
		newVersion.Left = vt.Tail
	}
	vt.Tail = newVersion
	return vt.Tail
}

func (vt *DocVersionTrack) RemoveHead() {
	if vt.Head == nil {
		return
	}

	vt.Head = vt.Head.Right

	if vt.Head != nil {
		vt.Head.Left = nil
	} else {
		vt.Tail = nil
	}
}

func (vt *DocVersionTrack) Count() int {
	count := 0
	node := vt.Head
	for node != nil {
		count++
		node = node.Right
	}
	return count
}
