package main

import (
	"fmt"
	"handsOn/internal/function"
	"handsOn/internal/gateway"
	"handsOn/internal/services"
)

func main() {
	fmt.Println("Experience a seamless payment service")

	pg := gateway.NewPaymentGateway()

	paypal := services.NewPaypal("paypal-123", "paypal-123", "paypal-123")
	stripe := services.NewStripe("stripe-456", "stripe-456", "stripe-456")
	razorpay := services.NewRazorpay("Razorpay-789", "RazorPay-789", "razorpay-789")

	pg.RegisterProvider(paypal)
	pg.RegisterProvider(stripe)
	pg.RegisterProvider(razorpay)

	paymentResponse, err := pg.ProcessPayment("PayPal", 100.50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(paymentResponse)
	}

	refundResponse, err := pg.IssueRefund("Stripe", "txn-001")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(refundResponse)
	}

	fmt.Println("Extracting provider details:")
	function.ExtractProviderDetails(paypal)
	function.ExtractProviderDetails(stripe)
	function.ExtractProviderDetails(razorpay)
}
