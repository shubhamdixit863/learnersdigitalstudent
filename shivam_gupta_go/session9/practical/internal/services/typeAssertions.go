package services

import "fmt"

func Exatractdetails(provider Services) {
	switch p := provider.(type) {
	case *PayPal:
		fmt.Printf("Provider: PayPal ; ClientID: %s ; APIKey: %s ; MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	case *Stripe:
		fmt.Printf("Provider: Stripe ; ClientID : %s ; APIKey : %s; MerchantID : %s\n",p.ClientID,p.APIKey,p.MerchantID)
	case *Razorpay:
		fmt.Printf("Provider: Stripe ; ClientID : %s ; APIKey : %s; MerchantID : %s\n",p.ClientID,p.APIKey,p.MerchantID)
	default: fmt.Println("error")
	}
}

