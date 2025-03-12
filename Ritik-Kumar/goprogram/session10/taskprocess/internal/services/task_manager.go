package services

import (
	"fmt"
	"log"
	"os"
	"time"
)

type TaskManager struct {
	TaskList    []Task
	FailedTasks []error
}

func (tm *TaskManager) AddTask(task Task) {
	tm.TaskList = append(tm.TaskList, task)
}

func (tm *TaskManager) ExecuteTasks() {
	for _, task := range tm.TaskList {
		retries := 0
		for retries < 3 {
			err := task.Run()
			if err != nil {
				log.Printf("Error executing task: %v (Retry: %d)\n", err, retries+1)
				retries++
				time.Sleep(1 * time.Second) 
				continue
			}
			log.Println("Task completed successfully")
			break
		}
		if retries == 3 {
			log.Println("Task failed after 3 retries")
			tm.FailedTasks = append(tm.FailedTasks, fmt.Errorf("Task failed: %v", task))
		}
	}
	tm.SaveFailedTasks()
}

func (tm *TaskManager) SaveFailedTasks() {
	if len(tm.FailedTasks) == 0 {
		return
	}
	file, err := os.Create("failed_tasks.log")
	if err != nil {
		log.Println("Error creating failed tasks log file:", err)
		return
	}
	defer file.Close()

	for _, taskErr := range tm.FailedTasks {
		fmt.Fprintln(file, taskErr)
	}
	log.Println("Failed tasks saved to log file")
}
