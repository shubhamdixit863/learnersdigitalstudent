package versionmanager

import (
	docversiontrack "document_versioning_system/internal/version_manager/doc_version_track"
	versiontrack "document_versioning_system/internal/version_manager/versionTrack"
	"errors"
)

type VersionMan struct {
	list        *docversiontrack.DocVersionTrack
	current     *docversiontrack.Node
	undo        *versiontrack.VersionTrack
	redo        *versiontrack.VersionTrack
	maxVersions int
}

func NewVersionManager(maxVersions int) *VersionMan {
	return &VersionMan{
		list:        docversiontrack.NewDocVersionTrack(),
		undo:        versiontrack.NewVersionTrack(),
		redo:        versiontrack.NewVersionTrack(),
		maxVersions: maxVersions,
	}
}

func (vm *VersionMan) AddVersion(content string) {
	vm.redo.Clear()

	newNode := vm.list.Append(content)

	vm.current = newNode

	if vm.maxVersions > 0 && vm.list.Count() > vm.maxVersions {
		vm.list.RemoveHead()
	}
}

func (vm *VersionMan) Undo() (string, error) {
	if vm.current == nil || vm.current.Left == nil {
		return "", errors.New("no previous version to undo")
	}
	vm.undo.Push(vm.current)
	vm.current = vm.current.Left
	return vm.current.Value.(string), nil
}

func (vm *VersionMan) Redo() (string, error) {
	if vm.redo.IsEmpty() {
		return "", errors.New("no version to redo")
	}
	lastUndone, err := vm.redo.Pop()
	if err != nil {
		return "", err
	}
	vm.current = lastUndone.(*docversiontrack.Node)
	return vm.current.Value.(string), nil
}

func (vm *VersionMan) GetCurrentVersion() string {
	if vm.current == nil {
		return ""
	}
	return vm.current.Value.(string)
}
