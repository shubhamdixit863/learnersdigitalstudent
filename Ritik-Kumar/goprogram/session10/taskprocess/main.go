package main

import (
    "log"
    "taskprocess/internal/services"
    "taskprocess/internal/utils"
)

func main() {
    utils.InitLogger()

    log.Println("Task Processing System Started")
    tm := &services.TaskManager{}   //concrete type task interface ko jo satisfy kraga voh pass hoga
    tm.AddTask(&services.FileProcessingTask{FileName: "data.txt"})
    tm.AddTask(&services.APICallTask{URL: "https://api.example.com"})
    tm.AddTask(&services.DataValidationTask{Data: ""}) 
    tm.ExecuteTasks()
    log.Println("Task Processing System Completed")
}
