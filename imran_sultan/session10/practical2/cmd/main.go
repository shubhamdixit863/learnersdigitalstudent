package main

import (
	"fmt"
	"practical2/internal/service"
)

func main() {
	vm := service.NewVersionManager(5)

	vm.AddVersion("Version 1")
	vm.AddVersion("Version 2")
	vm.AddVersion("Version 3")
	vm.AddVersion("Version 4")
	vm.AddVersion("Version 5")

	fmt.Println("Undo:", vm.Undo())
	fmt.Println("Undo:", vm.Undo())

	fmt.Println("Redo:", vm.Redo())

}
