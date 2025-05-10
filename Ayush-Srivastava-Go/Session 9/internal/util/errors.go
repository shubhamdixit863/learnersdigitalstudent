package util

import "fmt"

type PaymentError struct {
	message  string
	provider string
}

func NewPaymentError(provider, message string) *PaymentError {
	return &PaymentError{message: message, provider: provider}
}

func (pe *PaymentError) Error() string {
	return fmt.Sprintf("Payment Error: %s, Provider: %s", pe.provider, pe.message)
}

func (pe *PaymentError) PayPalError() string {
	return fmt.Sprintf("PayPal Error: %s", pe.message)
}

func (pe *PaymentError) StripeError() string {
	return fmt.Sprintf("Stripe Error: %s", pe.message)
}

func (pe *PaymentError) RazorPayError() string {
	return fmt.Sprintf("RazorPay Error: %s", pe.message)
}