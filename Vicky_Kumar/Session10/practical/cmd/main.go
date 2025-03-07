package main

import (
	"practical/internal/services"
)

func main() {
	// srvTsk := services.NewFileProcessing("./data/file.txt")
	// err := srvTsk.Run()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("File Processed Successfully")

	// srvTsk := services.NewDataValidation("vicky@gmail.com")
	// err := srvTsk.Run()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("Data Validated")

	// srvTsk := services.NewAPICall("https://jsonplaceholder.typicode.com/todos/1")
	// err := srvTsk.Run()
	// if err != nil {
	// 	log.Println("api eerr", err)
	// 	return
	// }
	// fmt.Println("API called")

	tskManager := services.NewTaskManager()
	filePrsng := services.NewFileProcessing("./data/file.txt")
	tskManager.AddTask(filePrsng)
	dataVldn := services.NewDataValidation("vickygmail.com")
	tskManager.AddTask(dataVldn)
	apicall := services.NewAPICall("https://jsonplaceholder.typicode.com/todos/1")
	tskManager.AddTask(apicall)
	tskManager.ExecuteTasks()

}
