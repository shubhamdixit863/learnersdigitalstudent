package utils

import "fmt"

const (
	fileOpeningError = "error: Could not open file %s"
)

func FileOpeningError(filename string) error {

	return fmt.Errorf(fileOpeningError, filename)
}
