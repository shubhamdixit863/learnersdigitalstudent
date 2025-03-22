package services

import (
	"encoding/json"
	"log"
	"os"
	"task/internal/utils"
	"time"
)


type TaskManager struct {
	Tasks      []utils.Task
	retryLimit int
	FailedLog  []utils.Task
}

func (tm *TaskManager) SetRetryLimit(limit int) {
	tm.retryLimit = limit
}

func (tm *TaskManager) AddTask(task utils.Task) {
	tm.Tasks = append(tm.Tasks, task)
}

func (tm *TaskManager) ExecuteTask() {
	for _, task := range tm.Tasks {
		success := false

		for i := 0; i < tm.retryLimit; i++ {
			err := task.Run()
			if err == nil {
				log.Printf("Task executed successfully")
				success = true
				break
			}
			log.Printf("Task failed at attemopt %d, error: %s", i, err.Error())
			time.Sleep(1 * time.Second)
		}

		if !success {
			tm.FailedLog = append(tm.FailedLog, task)
		}
	}
}


func (tm *TaskManager) SaveFailedTasks(filename string) error {

	file, err := os.Create(filename)

	if err != nil {
		log.Printf("Error creating file: %s", err.Error())
		return err
	}

	defer file.Close()
	
	jsonData, _ := json.Marshal(tm.FailedLog)
	file.Write(jsonData)
	return nil
}

