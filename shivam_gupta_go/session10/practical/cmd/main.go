package main

import (
	"session10/practical/internal/services"
)

func main() {
	ts := services.NewTaskManager(3)

	ts.AddTask(&services.FileProcessing{Filepath: "../test.txt"})
	ts.AddTask(&services.DataValidate{Data: "shivam@gmail.com",
		Pattern: `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`})

	ts.AddTask(&services.ApiCall{})
	ts.ExecuteTask()
  
	ts.Report()
}