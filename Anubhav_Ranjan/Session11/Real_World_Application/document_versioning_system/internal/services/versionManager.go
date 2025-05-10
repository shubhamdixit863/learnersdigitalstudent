package services

type VersionManager struct {
	head         *Node
	tail         *Node
	current      *Node
	redoStack    Stack
	versionCount int
	maxVersions  int
}

func NewVersionManager(maxVersions int) *VersionManager {
	if maxVersions <= 0 {
		maxVersions = 5
	}
	return &VersionManager{
		maxVersions: maxVersions,
	}
}

func (vm *VersionManager) AddVersion(content string) {
	newNode := &Node{Content: content}

	if vm.head == nil {
		vm.head = newNode
		vm.tail = newNode
		vm.current = newNode
		vm.versionCount = 1
		return
	}

	if vm.current != vm.tail {
		vm.tail = vm.current
		vm.redoStack.Clear()
	}

	newNode.Prev = vm.current
	vm.current.Next = newNode
	vm.current = newNode
	vm.tail = newNode
	vm.versionCount++

	if vm.versionCount > vm.maxVersions {
		vm.head = vm.head.Next
		if vm.head != nil {
			vm.head.Prev = nil
		}
		vm.versionCount--
	}
}

func (vm *VersionManager) Undo() string {
	if vm.current == nil || vm.current.Prev == nil {
		return vm.GetCurrentVersion()
	}

	vm.redoStack.Push(vm.current)

	vm.current = vm.current.Prev
	return vm.current.Content
}

func (vm *VersionManager) Redo() string {
	if vm.redoStack.IsEmpty() {
		return vm.GetCurrentVersion()
	}

	nextNode, err := vm.redoStack.Pop()
	if err != nil {
		return vm.GetCurrentVersion()
	}

	vm.current = nextNode
	return vm.current.Content
}

func (vm *VersionManager) GetCurrentVersion() string {
	if vm.current == nil {
		return ""
	}
	return vm.current.Content
}
