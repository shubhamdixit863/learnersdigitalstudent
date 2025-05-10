package services

import (
	"fmt"
	"practical2/internal/doublylinkedlist"
	"practical2/internal/stack"
)

type VersionManager struct {
	Versions    doublylinkedlist.DoublyLinkedList
	RedoStack   stack.Stack
	CurrVersion int
	MaxVersions int
}

func NewVersionManager(maxVersion int) *VersionManager {
	return &VersionManager{
		Versions:    *doublylinkedlist.NewDoublyLinkedList(),
		RedoStack:   *stack.NewStack(),
		CurrVersion: 0,
		MaxVersions: maxVersion,
	}
}

func (v *VersionManager) AddVersion(content string) {
	v.Versions.Insert(content, v.CurrVersion)
	v.CurrVersion++
	v.RedoStack = *stack.NewStack()
	if v.Versions.Size > v.MaxVersions {
		delPos := 0
		v.Versions.Delete(delPos)
	}
	fmt.Println("Added version:", content)
}

func (v *VersionManager) GetCurrentVersion() string {

	if v.CurrVersion == 0 {
		return "No versions available"
	}
	return v.Versions.GetData(v.CurrVersion - 1)
}

func (v *VersionManager) Undo() string {
	if v.CurrVersion == 0 || v.CurrVersion == 1 {
		return "No previous version available"
	}
	v.RedoStack.Push(v.CurrVersion)
	v.CurrVersion--
	fmt.Println("Undo to: ", v.GetCurrentVersion())
	return v.GetCurrentVersion()
}
func (v *VersionManager) Redo() string {
	if v.RedoStack.Empty() {
		return "No redo available"
	}
	v.CurrVersion = v.RedoStack.Top()
	fmt.Println("Redo to", v.GetCurrentVersion())
	return v.GetCurrentVersion()
}
func (v *VersionManager) PrintVersions() {
	v.Versions.PrintList()
}
