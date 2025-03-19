package services

type VersionManager struct {
	Versions *DoublyLinkedList
	Stack    Stack
	Current  *Node
	limit    int
}

func NewVersionManager(limit int) *VersionManager {
	dll := NewDoublyLinkedList(limit)
	return &VersionManager{
		Versions: dll,
	}
}

func (vm *VersionManager) AddVersion(content string) {
	vm.Versions.Add(content)
	vm.Current = vm.Versions.Tail
}

func (vm *VersionManager) Undo() string {
	if vm.Current == nil {
		return vm.GetCurrentVersion()
	}
	vm.Stack.Push(vm.Current)

	temp := vm.Current
	vm.Current = vm.Current.Prev
	vm.Versions.Tail = vm.Current
	vm.Current.Next = nil
	temp.Prev = nil

	return vm.Current.Content
}

func (vm *VersionManager) Redo() string {
	if vm.Stack.IsEmpty() {
		return vm.GetCurrentVersion()
	}

	content := vm.Stack.Pop()

	vm.Current.Next = content
	content.Prev = vm.Current
	vm.Current = content
	vm.Versions.Tail = vm.Current

	return vm.Current.Content
}

func (vm *VersionManager) GetCurrentVersion() string {
	if vm.Current != nil {
		return vm.Current.Content
	}
	return "No Version Available"
}
