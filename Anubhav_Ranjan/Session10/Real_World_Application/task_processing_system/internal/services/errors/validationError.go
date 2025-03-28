package errors

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{Field: field, Message: message}
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("Validation failed for field '%s': %s", ve.Field, ve.Message)
}
