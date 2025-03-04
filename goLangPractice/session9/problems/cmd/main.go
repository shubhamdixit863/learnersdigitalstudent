package main

import (
	"fmt"
	"session9/problems/internal/services"
)

func main() {
	pg := services.NewPaymentGateway()

	var PaymentProcessor1 services.PaymentProcessor
	PaymentProcessor1 = services.NewPayPal("paypal123", "12345", "98765")

	var PaymentProcessor4 services.PaymentProcessor
	PaymentProcessor4 = services.NewPayPal("paypal123", "345", "0123")

	var PaymentProcessor2 services.PaymentProcessor
	PaymentProcessor2 = services.NewRazorpay("razorpay123", "12345", "98765")

	var PaymentProcessor3 services.PaymentProcessor
	PaymentProcessor3 = services.NewStripe("stripe123", "12345", "98765")

	pg.RegisterProvider(&services.PayPal{})
	pg.RegisterProvider(&services.Stripe{})
	pg.RegisterProvider(&services.Razorpay{})

	payment, _ := pg.ProcessPayment("Stripe", 200)
	fmt.Println(payment)

	refund, _ := pg.IssueRefund("PayPal", "txn123")
	fmt.Println(refund)

	services.ExtractProviderDetails(PaymentProcessor1)
	services.ExtractProviderDetails(PaymentProcessor2)
	services.ExtractProviderDetails(PaymentProcessor3)
	services.ExtractProviderDetails(PaymentProcessor4)
}
