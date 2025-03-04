package function

import (
	"fmt"
	"handsOn/internal/services"
)

func ExtractProviderDetails(provider services.PaymentProcessor) {
	switch p := provider.(type) {
	case services.Paypal:
		fmt.Printf("PayPal Client ID: %s, APIKey: %s, MerchantId: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	case services.Stripe:

		fmt.Printf("Stripe Client ID: %s, APIKey: %s, MerchantId: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	case services.Razorpay:

		fmt.Printf("Razorpay Client ID: %s, APIKey: %s, MerchantId: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	default:
		fmt.Println("Unknown provider")
	}
}
