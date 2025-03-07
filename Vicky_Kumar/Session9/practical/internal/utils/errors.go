package utils

import "fmt"

type PayPalError struct {
	Message string
}
type RazorPayError struct {
	Message string
}
type StripeError struct {
	Message string
}

func NewPayPalError(message string) *PayPalError {
	return &PayPalError{
		Message: message,
	}
}
func NewRazorPayError(message string) *RazorPayError {
	return &RazorPayError{
		Message: message,
	}
}
func NewStripeError(message string) *StripeError {
	return &StripeError{
		Message: message,
	}
}

func (p *PayPalError) Error() string {
	return fmt.Sprintf("PayPal Error: %s", p.Message)
}

func (r *RazorPayError) Error() string {
	return fmt.Sprintf("RazorPay Error: %s", r.Message)
}

func (s *StripeError) Error() string {
	return fmt.Sprintf("Stripe Error: %s", s.Message)
}

// create a custom error for different payments methods
// and return errors from methods, if the operation fails as the second argument
// internal/utils/errors.go // PayPal error
