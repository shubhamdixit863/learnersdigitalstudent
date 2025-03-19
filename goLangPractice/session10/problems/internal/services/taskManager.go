package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"session10/problems/internal/utils"
)

type TaskManager struct {
	tasks       []Task
	failedTasks []string
}

func NewTaskManager() *TaskManager {
	return &TaskManager{tasks: []Task{}, failedTasks: []string{}}
}

func (tm *TaskManager) AddTask(task Task) {
	tm.tasks = append(tm.tasks, task)
}

func (tm *TaskManager) ExecuteTasks() {
	for _, task := range tm.tasks {
		for i := 0; i < 3; i++ {
			err := task.Run()
			if err == nil {
				log.Printf("Task %T completed successfully", task)
				break
			} else {
				log.Printf("Task %T failed: %v", task, err)
				if errors.Is(err, utils.TaskError) {
					log.Printf("Specific TaskError detected for task %T", task)
				}
				if i == 2 {
					tm.failedTasks = append(tm.failedTasks, fmt.Sprintf("Task %T failed: %v", task, err))
					log.Printf("Task %T marked as failed after 3 retries", task)
				}
			}
		}
	}
	tm.saveFailedTasksToJSON("FailedTasks.json")
}

func (tm *TaskManager) saveFailedTasksToJSON(filename string) {
	data, e := json.Marshal(tm.failedTasks)
	if e != nil {
		fmt.Printf("%w: Error marshaling JSON: %v\n", utils.TaskError, e)
		return
	}

	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Printf("%w: Error writing file: %v\n", utils.TaskError, err)
		return
	}

	fmt.Println("File written successfully!")
}
