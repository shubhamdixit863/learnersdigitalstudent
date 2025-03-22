package main

import (
	"fmt"
	"payment-gateway/internal/services"
)

func main() {

	gateway := services.NewPaymentGateway()
	transaction := services.NewTransactionService()


	PayPal := services.PayPal{
		ClientID: "paypal-client-id",
		APIKey:  "paypal-api-key",
		MerchantID: "paypal-merchant-id",
	}

	Stripe := services.Stripe{
		ClientID: "stripe-client-id",
		APIKey: "stripe-api-key",
		MerchantID: "stripe-merchant-id",
	}

	RazorPay := services.RazorPay{
		APIKey: "razorpay-api",
		MerchantID: "razorpay-merchant-id",
		ClientID: "razorpay-client-id",
	}

	gateway.RegisterProvider(&PayPal)
	gateway.RegisterProvider(&Stripe)
	gateway.RegisterProvider(&RazorPay)

	transaction1ID := transaction.CreateTransaction("Payal", 100.0)
	payement, err := gateway.ProcessPayment("PayPal", 100.0)

	if err != nil {
		fmt.Println("This line has error")
		panic(err)
	}
	fmt.Println(payement)

	isTransaction1Completed := transaction.UpdateTransactionStatus(transaction1ID, "COMPLETED")
	if !isTransaction1Completed {
		fmt.Println("Transaction not found")
	}

	transaction2ID := transaction.CreateTransaction("Stripe", 200.0)
	refund, err := gateway.IssueRefund("Stripe", transaction2ID)

	if err != nil {
		panic(err)
	}
	fmt.Println(refund)

	isTransaction2Completed := transaction.UpdateTransactionStatus(transaction2ID, "COMPLETED")
	if !isTransaction2Completed {
		fmt.Println("Transaction not found")
	}		

	services.GetProviderDetails(&PayPal)
	services.GetProviderDetails(&Stripe)
	services.GetProviderDetails(&RazorPay)

	for id, transaction := range transaction.GetAllTransactions() {
		fmt.Println("Transaction ID:", id)
		fmt.Println("Provider:", transaction.Provider)
		fmt.Println("Amount:", transaction.Amount)
		fmt.Println("Status:", transaction.Status)
		fmt.Println()
	}
}