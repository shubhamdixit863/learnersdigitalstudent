package services

import "fmt"

type Stripe struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewStripe(clientID, apiKey, merchantID string) PaymentProcessor {
	return &Stripe{
		ClientID:   clientID,
		APIKey:     apiKey,
		MerchantID: merchantID,
	}
}

func (p Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f via  Stripe (ClientID: %s)", amount, p.ClientID)
}

func (p Stripe) Refund(transactionID string) string {
	return fmt.Sprintf("Refunded transaction %s via Stripe", transactionID)
}

func (p Stripe) GetProviderName() string {
	return "Stripe"
}
