package services

import "fmt"

type TaskManager struct {
	Tasks []TaskService
}

func (tm *TaskManager) AddTask(task TaskService) {
	fmt.Println("append Task from TaskManager")
	tm.Tasks = append(tm.Tasks, task)
}

func (tm *TaskManager) Run() {
	for _, task := range tm.Tasks {
		err := task.Run()
		if !(err == nil) {
			fmt.Println(err)
		}
	}
}
