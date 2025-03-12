package services

import "fmt"

type PayPal struct { 
	ClientID   string
	APIKey     string
	MerchantID string
}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("PayPal processed payment of $%.2f", amount)
}

func (p *PayPal) Refund(transactionID string) string {
	return fmt.Sprintf("PayPal refunded transaction %s", transactionID)
}

func (p *PayPal) GetProviderName() string {
	return "PayPal"
}
