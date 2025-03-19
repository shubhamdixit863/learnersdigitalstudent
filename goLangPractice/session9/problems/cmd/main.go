package main

import (
	"fmt"
	"session9/problems/internal/services"
)

func main() {
	pg := services.NewPaymentGateway()

	var PaymentProcessor1 services.PaymentProcessor = services.NewPayPal("paypal123", "12345", "98765")

	var PaymentProcessor4 services.PaymentProcessor = services.NewPayPal("paypal123", "345", "0123")

	var PaymentProcessor2 services.PaymentProcessor = services.NewRazorpay("razorpay123", "12345", "98765")

	var PaymentProcessor3 services.PaymentProcessor = services.NewStripe("stripe123", "12345", "98765")

	pg.RegisterProvider(&services.PayPal{})
	pg.RegisterProvider(&services.Stripe{})
	pg.RegisterProvider(&services.Razorpay{})

	payment, err := pg.ProcessPayment("ss", 200)
	fmt.Println(payment)
	if err != nil {
		fmt.Println(err)
	}

	refund, err := pg.IssueRefund("PayPal", "")
	fmt.Println(refund)
	if err != nil {
		fmt.Println(err)
	}

	services.ExtractProviderDetails(PaymentProcessor1)
	services.ExtractProviderDetails(PaymentProcessor2)
	services.ExtractProviderDetails(PaymentProcessor3)
	services.ExtractProviderDetails(PaymentProcessor4)
}
