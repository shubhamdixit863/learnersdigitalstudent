package manager

import (
	"fmt"
	"log"
	"os"
	"task_processing_system/internal/services/tasks"
	"time"
)

type TaskManager struct {
	Tasks       []tasks.Task
	RetryLimit  int
	FailedTasks []string
}

func NewTaskManager(retrylimit int) *TaskManager {
	return &TaskManager{Tasks: []tasks.Task{}, RetryLimit: retrylimit, FailedTasks: []string{}}
}

func (tm *TaskManager) RegisterTask(task tasks.Task) {
	tm.Tasks = append(tm.Tasks, task)
}

func (tm *TaskManager) ExecuteTasks() {
	for _, task := range tm.Tasks {
		fmt.Println("Executing:", task.Name())

		for attempt := 1; attempt <= tm.RetryLimit; attempt++ {
			err := task.Run()
			if err != nil {
				fmt.Printf("Attempt %d failed: %v\n", attempt, err)

				if attempt == tm.RetryLimit {
					fmt.Println("Task failed after max retries:", task.Name())
					tm.logFailedTask(task, err)
				} else {
					time.Sleep(1 * time.Second)
				}
			} else {
				fmt.Println("Task completed successfully:", task.Name())
				break
			}
		}
	}
}

func (tm *TaskManager) logFailedTask(task tasks.Task, err error) {
	tm.FailedTasks = append(tm.FailedTasks, task.Name())

	file, errr := os.OpenFile("logs/task_errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errr != nil {
		fmt.Println(errr)
	}
	fmt.Println(1)
	defer file.Close()

	logger := log.New(file, "Error: ", log.LstdFlags)
	logger.Printf("Task %s failed: %v\n", task.Name(), err)
}
