package services

import (
	"fmt"
)

type PayPal struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewPayPal(clientID string, APIKey string, merchantID string) PaymentProcessor {
	return &PayPal{
		ClientID:   clientID,
		APIKey:     APIKey,
		MerchantID: merchantID,
	}
}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("PayPal: Processed payment of %v", amount)
}

func (p *PayPal) Refund(transactionID string) string {
	return fmt.Sprintf("PayPal: Refunded transaction %s", transactionID)
}

func (p *PayPal) GetProviderName() string {
	return "PayPal"
}
