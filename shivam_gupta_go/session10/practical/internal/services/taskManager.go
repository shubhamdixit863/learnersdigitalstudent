package services

import (
	"errors"
	"fmt"
	"log"
	"session10/practical/internal/utils"
)

type TaskManager struct {
	Tasks       []Task
	Failedtasks []string
	Limit       int
}

func NewTaskManager(limit int) *TaskManager {
	return &TaskManager{
		Tasks:       []Task{},
		Failedtasks: []string{},
		Limit:       limit,
	}
}

//add task

func (t *TaskManager) AddTask(task Task) {
	t.Tasks = append(t.Tasks, task)
}

//run task

var Valerr *utils.ValidateError
var Taskerr *utils.TaskError

func (t *TaskManager) ExecuteTask() {
	for i := range t.Tasks {
		var err error
		for j := 0; j < t.Limit; j++ {
			err = t.Tasks[i].Run()
			if err == nil {
				break
			}
			// if customErr, ok := err.(*utils.ValidateError); ok {
			// 	log.Println("Validation Error:", customErr)
			// } else if customErr, ok := err.(*utils.TaskError); ok {
			// 	log.Println("Task Error:", customErr)
			// } else {
			// 	log.Println("General Error:", err)
			// }
			if errors.As(err, &Valerr) {
				fmt.Println("Validation Error:", Valerr)
			} else if errors.As(err, &Taskerr) {
				fmt.Println(" Task Error:", Taskerr)
			}else{ fmt.Println(err)}

			log.Println("Task failed attempt:", j+1)
		}
		if err != nil {
			t.Failedtasks = append(t.Failedtasks, fmt.Sprintln(" task ", i+1, "failed after ", t.Limit, "retry"))
		}
	}
}

//log failed task

func (t *TaskManager) Report() {
	if len(t.Failedtasks) > 0 {
		log.Println("failed task report")
		for i := range t.Failedtasks {
			fmt.Println(t.Failedtasks[i])
		}
	} else {
		log.Println("All task completed successfully!")
	}
}
