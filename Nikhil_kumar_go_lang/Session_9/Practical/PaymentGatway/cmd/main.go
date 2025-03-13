package main

import (
	"fmt"
	"payment_gateway/internal/payment"
)

func main() {
	gateway := payment.NewPaymentGateway()

	paypal := payment.PayPal{ClientID: "paypal123", APIKey: "paypal-secret", MerchantID: "paypal456"} //New
	stripe := payment.Stripe{ClientID: "stripe123", APIKey: "stripe-secret", MerchantID: "stripe456"}
	razorpay := payment.Razorpay{ClientID: "razorpay123", APIKey: "razorpay-secret", MerchantID: "razorpay456"}

	gateway.RegisterProvider(paypal)
	gateway.RegisterProvider(stripe)
	gateway.RegisterProvider(razorpay)

	paymentResponse, err := gateway.ProcessPayment("PayPal", 100.50)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(paymentResponse)
	}

	refundResponse, err := gateway.IssueRefund("Stripe", "txn_789")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(refundResponse)
	}

	payment.ExtractProviderDetails(paypal)
	payment.ExtractProviderDetails(stripe)
	payment.ExtractProviderDetails(razorpay)
}