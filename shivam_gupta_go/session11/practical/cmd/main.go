package main

import (
	"fmt"
	"session11/practical/internal/services"
)

func main() {
	vm := services.NewVersionManager(5)
	vm.AddVersion("version1")
	vm.AddVersion("version2")
	vm.AddVersion("version3")
	vm.AddVersion("version4")
	vm.AddVersion("version5")

	fmt.Println(vm.GetCurrentVersion())
	fmt.Println("Undo")
	fmt.Println(vm.Undo())
	fmt.Println("Redo")
	fmt.Println(vm.Redo())

	vm.AddVersion("version6")

}
