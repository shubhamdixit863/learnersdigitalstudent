package main

import (
	"fmt"
	"task_manager_sys/internal/services"
)

func main() {
	tm := &services.TaskManager{}

	fmt.Println("call from main")
	fmt.Println(tm)

	tm.AddTask(&services.FileProcessingTask{FileName: "../file.txt"})
	tm.AddTask(&services.DataValidationTask{
		Email:         "abc@abc.com",
		Phone:         "1234567890",
		Age:           "23",
		Pattern_Email: "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$",
		Pattern_Phone: "^[0-9]{10}$",
		Pattern_Age:   "^[0-9]{2}$"})
	tm.AddTask(&services.APICallTask{})

	tm.Run()

}
