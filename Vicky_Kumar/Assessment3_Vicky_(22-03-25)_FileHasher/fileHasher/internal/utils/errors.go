package utils

const (
	ErrEmptyDirPath  = "empty directory path provided"
	ErrReadingDir    = "error reading directory: "
	ErrReadingFile   = "error reading file: "
	ErrEmptyFilePath = "empty file path provided"
	ErrNoFilesFound  = "no files found"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(message string) *CustomError {
	return &CustomError{
		Message: message,
	}
}
