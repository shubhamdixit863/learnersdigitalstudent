package services

import "fmt"

// . Implement Payment Providers

// • Create three struct implementations: PayPal, Stripe, and Razorpay.

// • Each struct should:

// • Store unique details (e.g., ClientID, APIKey, MerchantID).

// • Implement the PaymentProcessor interface


type Stripe struct{
    ClientID string 
		APIKey string 
		MerchantID string 
}

func (p *Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Stripe (clientId : %s)",amount,p.ClientID)
}

func (p *Stripe) Refund (transactionID string) string {
	return fmt.Sprintf("refund issued for transaction %s via Stripe",transactionID)
}

func (p *Stripe) GetProviderName() string {
	return "Stripe"
}