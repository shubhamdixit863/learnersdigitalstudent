package versionmanager

import (
	"session11/models"
	"session11/stack"
)

type VersionManager struct {
	head       *models.Node 
	tail       *models.Node 
	current    *models.Node 
	redoStack  stack.Stack  
	versionCap int          
	versionCnt int          
}

func NewVersionManager(capacity int) *VersionManager {
	return &VersionManager{versionCap: capacity}
}

func (vm *VersionManager) AddVersion(content string) {
	newNode := &models.Node{Content: content}

	vm.redoStack.Clear()

	if vm.current == nil {
		vm.head = newNode
		vm.tail = newNode
		vm.current = newNode
		vm.versionCnt = 1
		return
	}

	vm.tail.Next = newNode
	newNode.Prev = vm.tail
	vm.tail = newNode
	vm.current = newNode

	vm.versionCnt++
	if vm.versionCnt > vm.versionCap {
		vm.head = vm.head.Next
		vm.head.Prev = nil
		vm.versionCnt--
	}
}

func (vm *VersionManager) Undo() string {
	if vm.current == nil || vm.current.Prev == nil {
		return "No previous version available"
	}

	vm.redoStack.Push(vm.current)
	vm.current = vm.current.Prev

	return vm.current.Content
}

func (vm *VersionManager) Redo() string {
	if vm.redoStack.IsEmpty() {
		return "No redo available"
	}

	vm.current = vm.redoStack.Pop()

	return vm.current.Content
}

func (vm *VersionManager) GetCurrentVersion() string {
	if vm.current == nil {
		return "No version available"
	}
	return vm.current.Content
}
