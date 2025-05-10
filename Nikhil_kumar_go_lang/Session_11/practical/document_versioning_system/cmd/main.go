package main

import (
	versionmanager "document_versioning_system/internal/version_manager"
	"fmt"
)

func main() {
	vm := versionmanager.NewVersionManager(5)

	vm.AddVersion("Version 1")
	vm.AddVersion("Version 2")
	vm.AddVersion("Version 3")
	fmt.Println("Current Version:", vm.GetCurrentVersion()) 

	// Undo
	content, err := vm.Undo()
	if err != nil {
		fmt.Println("Undo Error:", err)
	} else {
		fmt.Println("After Undo:", content) 
	}

	// Undo again
	content, err = vm.Undo()
	if err != nil {
		fmt.Println("Undo Error:", err)
	} else {
		fmt.Println("After Undo:", content)
	}

	// Redo
	content, err = vm.Redo()
	if err != nil {
		fmt.Println("Redo Error:", err)
	} else {
		fmt.Println("After Redo:", content) 
	}

	// Add a new version after undoing
	vm.AddVersion("Version 4")
	fmt.Println("Current Version:", vm.GetCurrentVersion()) 

	// Undo
	content, err = vm.Undo()
	if err != nil {
		fmt.Println("Undo Error:", err)
	} else {
		fmt.Println("After Undo:", content) 
	}
}
