package versionManager

type Version struct {
	content string
	prev    *Version
	next    *Version
}

type VersionManager struct {
	head        *Version
	current     *Version
	redoStack   []string
	maxVersions int
	size        int
}

func NewVersionManager(maxVersions int) *VersionManager {
	return &VersionManager{
		head:        nil,
		current:     nil,
		redoStack:   []string{}, //
		maxVersions: maxVersions,
		size:        0,
	}
}

func (vm *VersionManager) AddVersion(content string) {
	if vm.current != nil {
		vm.redoStack = []string{}
	}

	newVersion := &Version{content: content, prev: vm.current}
	if vm.current != nil {
		vm.current.next = newVersion
	}
	vm.current = newVersion

	if vm.head == nil {
		vm.head = newVersion
	}

	vm.size++
	if vm.size > vm.maxVersions {
		vm.head = vm.head.next
		vm.head.prev = nil
		vm.size--
	}
}

func (vm *VersionManager) Undo() string {

	if vm.current != nil && vm.current.prev != nil {
		vm.redoStack = append(vm.redoStack, vm.current.content)
		vm.current = vm.current.prev
	}
	return vm.GetCurrentVersion()
}

func (vm *VersionManager) Redo() string {
	if len(vm.redoStack) > 0 && vm.current.next != nil {
		vm.current = vm.current.next
		vm.redoStack = vm.redoStack[:len(vm.redoStack)-1]
	}
	return vm.GetCurrentVersion()
}

func (vm *VersionManager) GetCurrentVersion() string {
	if vm.current != nil {
		return vm.current.content
	}
	return ""
}
