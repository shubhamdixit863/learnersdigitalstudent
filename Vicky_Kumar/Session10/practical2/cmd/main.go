package main

import (
	"fmt"
	"practical2/internal/services"
)

func main() {
	vm := services.NewVersionManager(5)

	vm.AddVersion("Version 1")
	vm.AddVersion("Version 2")
	vm.AddVersion("Version 3")
	fmt.Println("Current Version:", vm.GetCurrentVersion())

	vm.Undo()
	fmt.Println("Current Version:", vm.GetCurrentVersion())

	vm.Undo()
	fmt.Println("Current Version:", vm.GetCurrentVersion())

	vm.Redo()
	fmt.Println("Current Version:", vm.GetCurrentVersion())

	vm.AddVersion("Version 4")
	fmt.Println("Current Version:", vm.GetCurrentVersion())
	vm.PrintVersions()

	vm.Undo()
	fmt.Println("Current Version:", vm.GetCurrentVersion())

	vm.AddVersion("Version 5")

	vm.PrintVersions()

}
