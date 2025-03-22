package errorHandling

import "fmt"

type FileProcessingTaskError struct {
	msg string
}

func (f *FileProcessingTaskError) Error() string {
	return fmt.Sprintf("FileProcessingError: %s", f.msg)
}

type DataValidationTaskError struct {
	msg string
}

func (d *DataValidationTaskError) Error() string {
	return fmt.Sprintf("DataValidationError: %s", d.msg)
}

type APICallTask struct {
	msg string
}

func (a *APICallTask) Error() string {
	return fmt.Sprintf("ApiCallTask : %s", a.msg)
}
