package main

import (
	"fmt"
	"handsOn/internal/tasks"
	"handsOn/internal/util"
)

func main() {

	ts := util.NewTaskManager()
	fmt.Println("Welcome to task management system")
	filename := "test.txt"
	text := "My name is Go !"
	email := "ayush@gmail.com"
	phone := "995136"
	api := "https://example.com/api"

	t1 := tasks.NewFileProcessingTask(filename, text)
	t2 := tasks.NewDataValidationTask(email, phone)
	t3 := tasks.NewAPICallTask(api)
	ts.AddTask(t1)
	ts.AddTask(t2)
	ts.AddTask(t3)

	ts.RunTasks()
	ts.LogFailedTasks()
}
