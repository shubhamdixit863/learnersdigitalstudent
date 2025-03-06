package services

import (
	"log"
	"practical2/internal/utils"
)

type TaskManager struct {
	tasks     []Task
	taskerror []utils.TaskError
}

func (s *TaskManager) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

func (s *TaskManager) ExecuteTasks() {
	for _, tasks := range s.tasks {
		s.RetryTasks(tasks)
	}
}
func (s *TaskManager) RetryTasks(tasks Task) {
	for i := 0; i < 3; i++ {
		err := tasks.Run()
		if err == nil {
			log.Println("task successful")
			return
		}
		log.Println(err)
		log.Println(utils.NewTaskError())
	}
	log.Println("task failed 3 times")
	s.taskerror = append(s.taskerror, utils.TaskError{Taskerror: tasks.ReturnName()})

}
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}
func (s *TaskManager) GetErrorTasks() []utils.TaskError {
	return s.taskerror
}
