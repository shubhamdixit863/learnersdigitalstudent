package main

import (
	"fmt"
	"practical/internal/gateway"
	"practical/internal/services"
	"practical/internal/transactions"
	"practical/internal/utils"
)

func main() {

	var PaymentGateway = *gateway.NewPaymentGateway()

	paypalAPI := utils.RandomID()
	paypal := services.NewPayPalService(paypalAPI)
	PaymentGateway.RegisterProvider(paypal)

	razorpayAPI := utils.RandomID()
	razorpay := services.NewRazorPayService(razorpayAPI)
	PaymentGateway.RegisterProvider(razorpay)

	stripeAPI := utils.RandomID()
	stripe := services.NewStripeService(stripeAPI)
	PaymentGateway.RegisterProvider(stripe)

	detail, _ := PaymentGateway.ProcessPayment("PayPal", 123.5)
	fmt.Println(detail)
	txnID := "12345"
	detail2, _ := PaymentGateway.IssueRefund("PayPal", txnID)
	fmt.Println(detail2)

	PaymentGateway.ProviderDetail(paypal)
	PaymentGateway.ProviderDetail(stripe)

	transactions.ShowTransaction()
}
