package main

import (
	"fmt"
	"session11/version_manager"
)

func main() {
	vm := versionmanager.NewVersionManager(5) 

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

	vm.AddVersion("Version 5")
	vm.AddVersion("Version 6")
	vm.AddVersion("Version 7") 
	fmt.Println(vm.GetCurrentVersion()) 

	vm.Undo()
	vm.Undo()
	vm.Undo()
	vm.Undo() 
	fmt.Println(vm.GetCurrentVersion()) 
}
