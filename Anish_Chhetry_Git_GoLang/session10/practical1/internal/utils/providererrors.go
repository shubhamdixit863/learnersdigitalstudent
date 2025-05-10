package utils

// PROVIDER NOT FOUND ERROR
type ProviderNotFoundError struct {
}

func (c *ProviderNotFoundError) Error() string {
	return "Provider Not Found"
}
func NewProviderNotFoundError() error {
	return &ProviderNotFoundError{}
}
