package main

import (
	"fmt"
	"handsOn/internal/function"
	"handsOn/internal/gateway"
	"handsOn/internal/services"
)

func main() {
	fmt.Println("Experience a seamless payment service")
	clientid := "paypal-123"
	apikey := "paypal-123"
	merchantid := "paypal-123"

	pg := gateway.NewPaymentGateway()

	provider := services.NewPaypal(clientid, apikey, merchantid)
	//provider1 := services.NewStripe(clientid, apikey, merchantid)
	//razorpay := services.NewRazorpay("Razorpay-789", "RazorPay-789", "razorpay-789")

	pg.RegisterProvider(provider)
	//pg.RegisterProvider(provider1)
	//pg.RegisterProvider(razorpay)

	paymentResponse, err := pg.ProcessPayment("PayPal", 100.50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(paymentResponse)
	}

	refundResponse, err := pg.IssueRefund("PayPal", "txn-001")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(refundResponse)
	}

	fmt.Println("Extracting provider details:")
	function.ExtractProviderDetails(provider)
	//function.ExtractProviderDetails(stripe)
	//function.ExtractProviderDetails(razorpay)
}
