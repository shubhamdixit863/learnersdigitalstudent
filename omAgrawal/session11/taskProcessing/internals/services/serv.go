package services

type Task interface {
	Run() error
}

type FileProcessingTask struct {
}

func (f *FileProcessingTask) Run() error {
	return nil
}

type DataValidationTask struct {
}

func (d *DataValidationTask) Run() error {
	return nil
}

type APICallTask struct {
}

func (a *APICallTask) Run() error {
	return nil
}
