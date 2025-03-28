package main

import (
	"document_versioning_system/internal/services"
	"fmt"
)

func main() {
	vm := services.NewVersionManager(5)

	// Adding Versions
	vm.AddVersion("Version 1")
	vm.AddVersion("Version 2")
	vm.AddVersion("Version 3")
	fmt.Println(vm.GetCurrentVersion())

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion())

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion())

	vm.Redo()
	fmt.Println(vm.GetCurrentVersion())

	vm.AddVersion("Version 4")
	fmt.Println(vm.GetCurrentVersion())

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion())
}
