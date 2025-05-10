package main

import (
	"fmt"
	"task_processing_system/internal/services/manager"
	"task_processing_system/internal/services/tasks"
)

func main() {
	tm := manager.NewTaskManager(3)

	tm.RegisterTask(tasks.NewFileProcessingTask("../helloWorld.txt"))
	tm.RegisterTask(tasks.NewDataValidationTask("Anubhav", "anubhav@gmail.com", "9549936484"))
	tm.RegisterTask(tasks.NewAPICallTask("https://github.com/feelinganubhav"))
	tm.RegisterTask(tasks.NewFileProcessingTask("../hello.txt"))
	tm.RegisterTask(tasks.NewDataValidationTask("Anubhav12", "anubhav@gmail.com", "9549936484"))
	tm.RegisterTask(tasks.NewDataValidationTask("Anubhav", "anubhavgmail.com", "9549936484"))
	tm.RegisterTask(tasks.NewDataValidationTask("Anubhav", "anubhav@gmail.com", "9549936ab484"))
	tm.RegisterTask(tasks.NewAPICallTask("https://githhhub.com/feelinganubhav"))

	fmt.Println("Starting task execution...")
	tm.ExecuteTasks()


	if len(tm.FailedTasks) > 0 {
		fmt.Println("\nFailed Tasks:")
		for _, failedTask := range tm.FailedTasks {
			fmt.Println("-", failedTask)
		}
	}

}
