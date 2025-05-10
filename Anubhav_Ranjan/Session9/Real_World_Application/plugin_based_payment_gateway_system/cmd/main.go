package main

import (
	"fmt"
	"plugin_based_payment_gateway_system/internal/paymentProviders"
	"plugin_based_payment_gateway_system/internal/services"
)

func main() {
	gateway := services.NewPaymentGateway()

	paypal := paymentProviders.NewPayPal()
	stripe := paymentProviders.NewStripe()
	razorpay := paymentProviders.NewRazorPay()

	gateway.RegisterProvider(paypal)
	gateway.RegisterProvider(stripe)
	gateway.RegisterProvider(razorpay)

	for {
		fmt.Println("\n==== Payment Gateway System ====")
		fmt.Println("1. Process Payment")
		fmt.Println("2. Issue Refund")
		fmt.Println("3. Display Provider Details")
		fmt.Println("4. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var providerName string
			var amount float64
			fmt.Print("Enter Provider Name (PayPal, Stripe, Razorpay): ")
			fmt.Scan(&providerName)
			fmt.Print("Enter Amount: ")
			fmt.Scan(&amount)
			response, err := gateway.ProcessPayment(providerName, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(response)
			}

		case 2:
			var providerName, transactionID string
			fmt.Print("Enter Provider Name (PayPal, Stripe, Razorpay): ")
			fmt.Scan(&providerName)
			fmt.Print("Enter Transaction ID: ")
			fmt.Scan(&transactionID)
			response, err := gateway.IssueRefund(transactionID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(response)
			}

		case 3:
			var providerName string
			fmt.Print("Enter Provider Name (PayPal, Stripe, Razorpay): ")
			fmt.Scan(&providerName)
			gateway.DisplayProviderDetails(providerName)

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
