package main

import (
	"fmt"
	"payment_gateway/internal/services"
)

const (
	PayPal_ClientID   = "paypal123"
	PayPal_APIKey     = "paypal-secret"
	PayPal_MerchantID = "paypal456"

	Stripe_ClientID   = "stripe123"
	Stripe_APIKey     = "stripe-secret"
	Stripe_MerchantID = "stripe456"

	Razorpay_ClientID   = "razorpay123"
	Razorpay_APIKey     = "razorpay-secret"
	Razorpay_MerchantID = "razorpay456"
)

func main() {
	gateway := services.NewPaymentGateway()
	paypal := services.PayPal{
		ClientID:   PayPal_ClientID,
		APIKey:     PayPal_APIKey,
		MerchantID: PayPal_MerchantID,
	}
	stripe := services.Stripe{
		ClientID:   Stripe_ClientID,
		APIKey:     Stripe_APIKey,
		MerchantID: Stripe_MerchantID,
	}
	razorpay := services.Razorpay{
		ClientID:   Razorpay_ClientID,
		APIKey:     Razorpay_APIKey,
		MerchantID: Razorpay_MerchantID,
	}

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

	services.ExtractProviderDetails(paypal)
	services.ExtractProviderDetails(stripe)
	services.ExtractProviderDetails(razorpay)
}
