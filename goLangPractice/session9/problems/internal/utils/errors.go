package utils

import (
	"fmt"
)

type PaymentGatewayError struct {
	Msg          string
	ProviderName string
}

func NewGatewayError(msg string, providerName string) *PaymentGatewayError {
	return &PaymentGatewayError{Msg: msg, ProviderName: providerName}
}

func (ge *PaymentGatewayError) Error() string {
	switch ge.ProviderName {
	case "PayPal":
		customError := NewPayPalError(ge.Msg)
		return customError.Error()
	case "Stripe":
		customError := NewStripeError(ge.Msg)
		return customError.Error()
	case "Razorpay":
		customError := NewRazorPayError(ge.Msg)
		return customError.Error()
	default:
		customError := fmt.Sprintln("Unknown provider", ge.ProviderName)
		return customError
	}
}

type PayPalError struct {
	Msg string
}

func NewPayPalError(msg string) *PayPalError {
	return &PayPalError{Msg: msg}
}

func (p *PayPalError) Error() string {
	return fmt.Sprintf("PayPal Error: %s", p.Msg)
}

type StripeError struct {
	Msg string
}

func NewStripeError(msg string) *StripeError {
	return &StripeError{Msg: msg}
}

func (s *StripeError) Error() string {
	return fmt.Sprintf("Stripe Error: %s", s.Msg)
}

type RazorPayError struct {
	Msg string
}

func NewRazorPayError(msg string) *RazorPayError {
	return &RazorPayError{Msg: msg}
}

func (r *RazorPayError) Error() string {
	return fmt.Sprintf("RazorPay Error: %s", r.Msg)
}
