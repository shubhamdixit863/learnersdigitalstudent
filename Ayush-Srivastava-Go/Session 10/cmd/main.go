package main

import (
	"log"
	"os"
	"task/internal/services"
)

func main() {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	defer logFile.Close()
	log.SetOutput(logFile)

	tm := services.TaskManager{}
	tm.SetRetryLimit(3)

	tm.AddTask(&services.FileProcessing{FileName: "text.txt"})
	tm.AddTask(&services.DataValidation{Data: ""})
	tm.AddTask(&services.APICall{Endpoint: "http://google.com"})

	tm.ExecuteTask()
	
}