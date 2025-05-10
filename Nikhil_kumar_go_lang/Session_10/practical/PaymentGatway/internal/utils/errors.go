package utils

import "fmt"

type Payment_Error struct {
	Msg    string
	P_name string
}

func (e *Payment_Error) Error() string {
	return fmt.Sprintf("Payment Error [%s]: %s", e.P_name, e.Msg)
}

func (e *Payment_Error) IsProviderError() bool {
	return e.P_name != ""
}

func NewPayment_Error(msg, provider string) *Payment_Error {
	return &Payment_Error{Msg: msg, P_name: provider}
}

type Paypal_Error struct {
	Msg    string
	P_name string
}

func (e *Paypal_Error) Error() string {
	return fmt.Sprintf("PayPal Error [%s]: %s", e.P_name, e.Msg)
}

func NewPaypal_Error(msg, provider string) *Paypal_Error {
	return &Paypal_Error{Msg: msg, P_name: provider}
}

type Stripe_Error struct {
	Msg    string
	P_name string
}

func (e *Stripe_Error) Error() string {
	return fmt.Sprintf("Stripe Error [%s]: %s", e.P_name, e.Msg)
}

func NewStripe_Error(msg, provider string) *Stripe_Error {
	return &Stripe_Error{Msg: msg, P_name: provider}
}

type Razorpay_Error struct {
	Msg    string
	P_name string
}

func (e *Razorpay_Error) Error() string {
	return fmt.Sprintf("Razorpay Error [%s]: %s", e.P_name, e.Msg)
}

func NewRazorpay_Error(msg, provider string) *Razorpay_Error {
	return &Razorpay_Error{Msg: msg, P_name: provider}
}
