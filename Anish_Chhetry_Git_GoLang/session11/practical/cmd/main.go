package main

import (
	"log"
	"practical/internal/services"
)

func main() {
	a := services.NewVersionManager(5)
	a.AddVersion("Version 1")
	a.AddVersion("Version 2")
	a.AddVersion("Version 3")
	log.Println(a.GetCurrentVersion())
	a.Undo()
	log.Println(a.GetCurrentVersion())
	a.Redo()
	log.Println(a.GetCurrentVersion())
}
