package utils

type ValidationError struct {
}

func (c *ValidationError) Error() string {
	return "Validation error"
}
func NewValidationError() *ValidationError {
	return &ValidationError{}
}
