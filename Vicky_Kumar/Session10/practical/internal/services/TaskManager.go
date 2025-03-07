package services

import (
	"log"
	"practical/internal/errorss"
)

type TaskManager struct {
	Tasks       []Task
	FailedTasks []error
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}
func (tm *TaskManager) AddTask(t Task) {
	tm.Tasks = append(tm.Tasks, t)
}

func (tm *TaskManager) ExecuteTasks() {
	for _, task := range tm.Tasks {
		retries := 0
		for retries < 3 {
			err := task.Run()
			if err == nil {
				log.Println("Task executed successfully!")
				break
			}
			retries++
			// log.Printf("Task failed, (attempt %d): %v\n", retries, err)
			errorss.HandleError(err)

			if retries == 3 {
				tm.FailedTasks = append(tm.FailedTasks, err)
			}
		}
	}

}
