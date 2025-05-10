package tasks

type Task interface {
	Run() error
	Name() string
}
