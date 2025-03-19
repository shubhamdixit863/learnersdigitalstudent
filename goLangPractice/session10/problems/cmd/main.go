package main

import "session10/problems/internal/services"

func main() {
	taskManager := services.NewTaskManager()

	dataValidationTask := services.NewDataValidationTask("Sagar Sinha", "sinhasagar89@gamil.com", "0123456789")
	fileProcessingTask := services.NewFileProcessingTask("HelloWorld.txt")
	apiCallTask := services.NewAPICallTask("")

	taskManager.AddTask(fileProcessingTask)
	taskManager.AddTask(dataValidationTask)
	taskManager.AddTask(apiCallTask)

	taskManager.ExecuteTasks()
}
