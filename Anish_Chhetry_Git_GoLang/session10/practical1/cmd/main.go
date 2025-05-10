package main

import (
	"fmt"
	"practical/internal/services"
)

func main() {
	gateway := services.NewPaymentGateway()
	paypal := services.NewPaypal("PC1", "PK1", "PM1")
	gateway.RegisterProvider(paypal)

	PpaymentResponse, PerrP := gateway.ProcessPayment("PayPal", 0)
	if PerrP != nil {
		fmt.Println(PerrP)
	}
	fmt.Println(PpaymentResponse)

	PrefundResponse, PerrR := gateway.IssueRefund("PayPal", "")
	if PerrR != nil {
		fmt.Println(PerrR)
	}
	fmt.Println(PrefundResponse)

	Extractor1, EErr1 := services.ExtractProvider(paypal)
	if EErr1 != nil {
		fmt.Println(EErr1)
	}
	fmt.Println(Extractor1)
	fmt.Println("-----------------------------------------------")
	stripe := services.NewStripe("SC1", "SK1", "SM1")
	gateway.RegisterProvider(stripe)

	SpaymentResponse, SerrP := gateway.ProcessPayment("Stripe", -33)
	if SerrP != nil {
		fmt.Println(SerrP)
	}
	fmt.Println(SpaymentResponse)

	SrefundResponse, SerrR := gateway.IssueRefund("Stripe", "")
	if SerrR != nil {
		fmt.Println(SerrR)
	}
	fmt.Println(SrefundResponse)

	Extractor2, EErr2 := services.ExtractProvider(stripe)
	if EErr2 != nil {
		fmt.Println(EErr2)
	}
	fmt.Println(Extractor2)
	fmt.Println("-----------------------------------------------")

	razorpay := services.NewRazorpay("RC1", "RK1", "RM1")
	gateway.RegisterProvider(razorpay)

	RpaymentResponse, RerrP := gateway.ProcessPayment("Razorpay", -12)
	if RerrP != nil {
		fmt.Println(RerrP)
	}
	fmt.Println(RpaymentResponse)

	RrefundResponse, RerrR := gateway.IssueRefund("Razorpays", "")
	if RerrR != nil {
		fmt.Println(RerrR)
	}
	fmt.Println(RrefundResponse)

	Extractor3, EErr3 := services.ExtractProvider(razorpay)
	if EErr3 != nil {
		fmt.Println(EErr3)
	}
	fmt.Println(Extractor3)
}
