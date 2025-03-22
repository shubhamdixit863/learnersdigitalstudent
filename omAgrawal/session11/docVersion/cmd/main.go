package main

import "docVersion/internals/services"

func main() {

	vm := services.NewVm()
	vm.AddVersion("version 1")
	vm.AddVersion("version2")
	vm.AddVersion("version 3")
	vm.UNDO()
	vm.GetCurrentversion()
	//vm.Redo()

}
