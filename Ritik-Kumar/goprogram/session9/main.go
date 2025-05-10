package main

import (
	"fmt"
	"session9/internal/services"
)

func main() {
	gateway := services.NewPaymentGateway()

	paypal:=&services.PayPal{ClientID: "paypal123",APIKey: "123API",MerchantID: "SGSDSJD"}

	gateway.RegisterProvider(paypal)
	
	paymentMsg, err := gateway.ProcessPayment("PayPal", 150.75)
	if err != nil {
		fmt.Println("Payment failed:", err)
	} else {
		fmt.Println(paymentMsg)
	}

	refundMsg, err := gateway.IssueRefund("PayPal", "txn_98765")
	if err != nil {
		fmt.Println("Refund failed:", err)
	} else {
		fmt.Println(refundMsg)
	}
	pro:=gateway.Providers["PayPal"]

	services.GetProviderDetails(pro)
}




//Create a custom error for different payments method
//and return errors from methods ,if the operations fails as the second argument 

//internal/service
//internal/utils/error.go  //paypal error