package utils

type TaskError struct {
	Taskerror string
}

func (c *TaskError) Error() string {
	return "Task error, Retrying again"
}
func NewTaskError() *TaskError {
	return &TaskError{}
}
