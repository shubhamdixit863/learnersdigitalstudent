package main

import (
	"fmt"
	"versionM/internal/services"
)

func main() {

	vm := services.NewVersionManager(5)

	vm.AddVersion("v1")
	vm.AddVersion("v2")
	vm.AddVersion("v3")
	vm.AddVersion("v4")

	fmt.Println(vm.GetCurrentVersion())

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion())

	vm.Undo()
	fmt.Println(vm.GetCurrentVersion())

	vm.Redo()
	fmt.Println(vm.GetCurrentVersion())

}
