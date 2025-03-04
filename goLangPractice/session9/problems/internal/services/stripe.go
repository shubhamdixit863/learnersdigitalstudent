package services

import (
	"fmt"
)

type Stripe struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewStripe(clientID string, APIKey string, merchantID string) PaymentProcessor {
	return &Stripe{
		ClientID:   clientID,
		APIKey:     APIKey,
		MerchantID: merchantID,
	}
}

func (p *Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Stripe: Processed payment of %v", amount)
}

func (p *Stripe) Refund(transactionID string) string {
	return fmt.Sprintf("Stripe: Refunded transaction %s", transactionID)
}

func (p *Stripe) GetProviderName() string {
	return "Stripe"
}
