package util

import (
	"fmt"
	"handsOn/internal/tasks"
	"log"
	"os"
	"time"
)

type TaskManager struct {
	Tasks       []tasks.Task
	FailedTasks []error
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks:       []tasks.Task{},
		FailedTasks: []error{},
	}
}

func (tm *TaskManager) AddTask(task tasks.Task) {
	tm.Tasks = append(tm.Tasks, task)
}

// RunTasks executes tasks sequentially with retries.
func (tm *TaskManager) RunTasks() {
	for _, task := range tm.Tasks {
		retries := 0
		for {
			err := task.Run()
			if err == nil {
				fmt.Println("Task succeeded")
				break
			}

			fmt.Printf("Task failed: %v\n", err)
			retries++

			if retries >= 3 {
				fmt.Println("Max retries reached. Marking task as failed.")
				tm.FailedTasks = append(tm.FailedTasks, err)
				break
			}

			fmt.Println("Retrying task...")
			time.Sleep(1 * time.Second) // Simulating retry delay
		}
	}
}

func (tm *TaskManager) LogFailedTasks() {
	file, err := os.Create("failed_tasks.log")
	if err != nil {
		log.Fatalf("Could not create log file: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	for _, err := range tm.FailedTasks {
		logger.Println(err)
	}
}
