package service

type Node struct {
	Data string
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
	Limit  int
}

type VersionManager struct {
	Versions  DoublyLinkedList
	RedoStack []*Node
	Current   *Node
}

func NewVersionManager(limit int) *VersionManager {
	return &VersionManager{
		Versions:  DoublyLinkedList{Limit: limit},
		RedoStack: []*Node{},
	}
}

func (vm *VersionManager) AddVersion(content string) {
	newNode := &Node{Data: content}

	if vm.Versions.Head == nil {
		vm.Versions.Head = newNode
		vm.Versions.Tail = newNode
		vm.Current = newNode
	} else {

		vm.Versions.Tail.Next = newNode
		newNode.Prev = vm.Versions.Tail
		vm.Versions.Tail = newNode
		vm.Current = newNode
	}

	vm.Versions.Length++
	if vm.Versions.Length > vm.Versions.Limit {

		vm.Versions.Head = vm.Versions.Head.Next
		if vm.Versions.Head != nil {
			vm.Versions.Head.Prev = nil
		}
		vm.Versions.Length--
	}

	vm.RedoStack = nil
}

func (vm *VersionManager) GetCurrentVersion() string {
	if vm.Current == nil {
		return "No version available"
	}
	return vm.Current.Data
}

func (vm *VersionManager) Undo() string {
	if vm.Current == nil || vm.Current.Prev == nil {
		return "No previous version available"
	}
	vm.RedoStack = append(vm.RedoStack, vm.Current)

	vm.Current = vm.Current.Prev
	return vm.Current.Data
}

func (vm *VersionManager) Redo() string {
	if len(vm.RedoStack) == 0 {
		return "No redo available"
	}

	redoVersion := vm.RedoStack[len(vm.RedoStack)-1]
	vm.RedoStack = vm.RedoStack[:len(vm.RedoStack)-1]
	vm.Current = redoVersion
	return vm.Current.Data
}
