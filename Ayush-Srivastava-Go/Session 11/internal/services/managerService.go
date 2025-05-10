package services

type VersionManager struct {
	curr      *Node
	redoStack []*Node
	capacity  int
	size      int
	head      *Node
}

func NewVersionManager(capacity int) *VersionManager {
	return &VersionManager{
		capacity: capacity,
	}
}

func (vm *VersionManager) AddVersion(versionName string) {

	newNode := &Node{data: versionName}
	if vm.curr == nil {
		vm.curr = newNode
		vm.head = newNode
		vm.size = 1
	} else {
		newNode.prev = vm.curr
		vm.curr.next = newNode
		vm.curr = newNode
		vm.size++
	}

	vm.redoStack = nil

	if vm.size > vm.capacity {
		vm.head = vm.head.next
		vm.head.prev = nil
		vm.size--
	}
}

func (vm *VersionManager) GetCurrentVersion() string {
	if vm.curr == nil {
		return ""
	}
	return vm.curr.data
}

func (vm *VersionManager) Undo() string {

	if vm.curr == nil || vm.curr.prev == nil {
		return vm.GetCurrentVersion()
	}

	vm.redoStack = append(vm.redoStack, vm.curr)
	vm.curr = vm.curr.prev
	return vm.curr.data
}

func (vm *VersionManager) Redo() string {

	if len(vm.redoStack) == 0 {
		return vm.GetCurrentVersion()
	}

	vm.curr = vm.redoStack[len(vm.redoStack)-1]
	vm.redoStack = vm.redoStack[:len(vm.redoStack)-1]
	return vm.curr.data
}
