package main

import (
	"fmt"
	"handsOn/internal/versionManager"
)

func main() {
	vm := versionManager.NewVersionManager(5)

	vm.AddVersion("Version 1")
	vm.AddVersion("Version 2")
	vm.AddVersion("Version 3")
	fmt.Println(vm.GetCurrentVersion()) // Output: "Version 3"

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion()) // Output: "Version 2"

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion()) // Output: "Version 1"

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion()) // Output: "Version 1"

	vm.Redo()
	fmt.Println(vm.GetCurrentVersion()) // Output: "Version 2"

	vm.AddVersion("Version 4")          // This clears the redo stack
	fmt.Println(vm.GetCurrentVersion()) // Output: "Version 4"

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion()) // Output: "Version 2"

}
