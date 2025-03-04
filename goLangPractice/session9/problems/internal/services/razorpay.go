package services

import (
	"fmt"
)

type Razorpay struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewRazorpay(clientID string, APIKey string, merchantID string) PaymentProcessor {
	return &Razorpay{
		ClientID:   clientID,
		APIKey:     APIKey,
		MerchantID: merchantID,
	}
}

func (p *Razorpay) Pay(amount float64) string {
	return fmt.Sprintf("Razorpay: Processed payment of %v", amount)
}

func (p *Razorpay) Refund(transactionID string) string {
	return fmt.Sprintf("Razorpay: Refunded transaction %s", transactionID)
}

func (p *Razorpay) GetProviderName() string {
	return "Razorpay"
}
