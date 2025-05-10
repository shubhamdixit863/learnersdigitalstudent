package main

import (
	"log"
	"practical2/internal/services"
)

func main() {
	s := services.NewTaskManager()
	s.AddTask(services.NewFileProcessingTask("/Users/anishchhetry/Documents/GoLang/session10/practical2/hello.txt"))
	s.AddTask(services.NewDataValidationTask("/Users/anishchhetry/Documents/GoLang/session10/practical2/hello.txt"))
	s.AddTask(services.NewAPICallTask("api.exe"))
	s.AddTask(services.NewFileProcessingTask("hello.txt"))
	s.AddTask(services.NewDataValidationTask("/Users/anishchhetry/Documents/GoLang/session10/practical2/hello2.txt"))
	s.AddTask(services.NewAPICallTask(""))
	s.ExecuteTasks()
	log.Println("Tasks where errors occurred: ", s.GetErrorTasks())
}
