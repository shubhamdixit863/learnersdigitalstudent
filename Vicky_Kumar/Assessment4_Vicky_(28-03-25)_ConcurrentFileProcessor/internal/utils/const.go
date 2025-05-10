package utils

const (
	ErrEmptyDirPath  = "empty directory path provided"
	ErrReadingDir    = "error reading directory: "
	ErrReadingFile   = "error reading file: "
	ErrEmptyFilePath = "empty file path provided"
	ErrNoFilesFound  = "no files found"
	ErrApiProcessing = "Error processing line '%s': %v"
	ErrApiRetry      = "Error Processing line %s, Retrying...\n"
	WrongChoice      = "Please enter a valid choice."

	WordCount   = "Total word count: %d"
	ApiResponse = "API Response for line '%s': %s"
	KeyWord     = "Enter the keyword to search"
	EmptyString = ""
)

var PrintOptions = []string{"1. Line Filter", "2. Word Count", "3. Api Call", "4. Retru Api Call"}
var Options = []string{"filter", "wordcount", "api", "retry"}

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
