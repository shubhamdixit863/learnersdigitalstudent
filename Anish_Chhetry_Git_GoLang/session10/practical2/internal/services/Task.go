package services

type Task interface {
	Run() error
	ReturnName() string
}
