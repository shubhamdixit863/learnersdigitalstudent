package customErrors

import (
	"errors"
	"fmt"
	"plugin_based_payment_gateway_system/internal/paymentInterface"
	"plugin_based_payment_gateway_system/internal/paymentProviders"
)

type GatewayError struct {
	Msg      string
	Provider paymentInterface.PaymentProcessor
}

func NewGatewayError(msg string, provider paymentInterface.PaymentProcessor) *GatewayError {
	return &GatewayError{Msg: msg, Provider: provider}
}

func (ge *GatewayError) Error() string {
	var customError error
	switch p := ge.Provider.(type) {
	case *paymentProviders.PayPal:
		customError = NewPayPalError(ge.Msg)

	case *paymentProviders.Stripe:
		customError = NewStripeError(ge.Msg)

	case *paymentProviders.RazorPay:
		customError = NewRazorPayError(ge.Msg)

	default:
		err := fmt.Sprintf("Unknown provider: %+v, Message: %s", p, ge.Msg)
		customError = errors.New(err)
	}

	return customError.Error()

}

type PayPalError struct {
	Msg string
}

func NewPayPalError(msg string) *PayPalError {
	return &PayPalError{Msg: msg}
}

func (pe *PayPalError) Error() string {
	return fmt.Sprintf("PayPal Error: %s", pe.Msg)
}

type StripeError struct {
	Msg string
}

func NewStripeError(msg string) *StripeError {
	return &StripeError{Msg: msg}
}

func (se *StripeError) Error() string {
	return fmt.Sprintf("Stripe Error: %s", se.Msg)
}

type RazorPayError struct {
	Msg string
}

func NewRazorPayError(msg string) *RazorPayError {
	return &RazorPayError{Msg: msg}
}

func (re *RazorPayError) Error() string {
	return fmt.Sprintf("RazorPay Error: %s", re.Msg)
}
