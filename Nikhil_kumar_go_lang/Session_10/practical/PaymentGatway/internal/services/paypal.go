package services

import "fmt"


type PayPal struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func (p PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal", amount)
}

func (p PayPal) Refund(transactionID string) string {
	return fmt.Sprintf("Refund issued for transaction %s via PayPal", transactionID)
}

func (p PayPal) GetProviderName() string {
	return "PayPal"
}
