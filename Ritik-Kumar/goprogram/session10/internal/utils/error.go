package utils

import "fmt"

type PaymentError struct {
	Provider string
	Message  string
} 

func (e *PaymentError) Error() string {
	return fmt.Sprintf("[%s Error] %s", e.Provider, e.Message)
}

func NewPaymentError(provider, message string) error {
	return &PaymentError{Provider: provider, Message: message}
}

type StripeError struct {
	Provider string
	Message  string
} 

func (e *StripeError) Error() string {
	return fmt.Sprintf("[%s Error] %s", e.Provider, e.Message)
}

func NewStripeError(provider, message string) error {
	return &StripeError{Provider: provider, Message: message}
}

