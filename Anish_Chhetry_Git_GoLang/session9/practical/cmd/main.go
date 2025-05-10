package main

import (
	"fmt"
	"practical/internal/services"
)

func main() {
	gateway := services.NewPaymentGateway()
	paypal := services.NewPaypal("PC1", "PK1", "PM1")
	gateway.RegisterProvider(paypal)

	PpaymentResponse, _ := gateway.ProcessPayment("PayPal", 100)
	fmt.Println(PpaymentResponse)

	PrefundResponse, _ := gateway.IssueRefund("PayPal", "ads12e1")
	fmt.Println(PrefundResponse)

	services.ExtractProvider(paypal)
	fmt.Println("-----------------------------------------------")
	stripe := services.NewStripe("SC1", "SK1", "SM1")
	gateway.RegisterProvider(stripe)

	SpaymentResponse, _ := gateway.ProcessPayment("Stripe", 100)
	fmt.Println(SpaymentResponse)

	SrefundResponse, _ := gateway.IssueRefund("Stripe", "19238dsa")
	fmt.Println(SrefundResponse)

	services.ExtractProvider(stripe)
	fmt.Println("-----------------------------------------------")

	razorpay := services.NewRazorpay("RC1", "RK1", "RM1")
	gateway.RegisterProvider(razorpay)

	RpaymentResponse, _ := gateway.ProcessPayment("Razorpay", 100)
	fmt.Println(RpaymentResponse)

	RrefundResponse, _ := gateway.IssueRefund("Razorpay", "adfk13")
	fmt.Println(RrefundResponse)

	services.ExtractProvider(razorpay)
}
