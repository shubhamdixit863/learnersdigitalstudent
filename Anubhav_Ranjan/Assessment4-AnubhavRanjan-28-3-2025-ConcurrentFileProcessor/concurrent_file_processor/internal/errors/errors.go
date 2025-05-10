package errors

import (
	"fmt"
)

const (
	ErrInvalidDirectoryMsg       = "invalid directory path: %v"
	ErrNoTextFilesMsg            = "no text files found in the directory"
	ErrProcessingFailedMsg       = "file processing failed: %v"
	ErrReadingFileMsg            = "error reading file %s: %v"
	ErrAPICallFailed             = "API call failed"
	ErrInvalidDirectory          = "invalid directory path"
	ErrNoTextFiles               = "no text files found in the directory"
	ErrProcessingFailed          = "file processing failed"
	ErrUnsupportedProcessingMode = "unsupported processing mode"
	ErrOpenFile                  = "failed to open file %s: %v"
	ErrProcessingFmt             = "Processing error: %v\n"
)

func LogError(err error) string {
	return fmt.Sprintf("Error: %+v", err)
}
