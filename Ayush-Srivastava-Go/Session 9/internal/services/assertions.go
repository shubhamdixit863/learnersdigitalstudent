package services

import (
	"fmt"
)

func GetProviderDetails(provider PaymentProcessor) {
	switch p := provider.(type) {
	case PayPal:
		fmt.Println("PayPal Client ID:", p.ClientID)
		fmt.Println("PayPal API Key:", p.APIKey)
		fmt.Println("PayPal Merchant ID:", p.MerchantID)
	case Stripe:
		fmt.Println("Stripe Client ID:", p.ClientID)
		fmt.Println("Stripe API Key:", p.APIKey)
		fmt.Println("Stripe Merchant ID:", p.MerchantID)
	case RazorPay:
		fmt.Println("RazorPay API Key:", p.APIKey)
		fmt.Println("RazorPay Merchant ID:", p.MerchantID)
		fmt.Println("RazorPay Client ID:", p.ClientID)
	}
}