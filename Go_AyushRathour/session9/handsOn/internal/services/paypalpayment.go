package services

import "fmt"

type Paypal struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewPaypal(clientID, apiKey, merchantID string) PaymentProcessor {
	return &Paypal{
		ClientID:   clientID,
		APIKey:     apiKey,
		MerchantID: merchantID,
	}
}
func (p Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f via PayPal (ClientID: %s)", amount, p.ClientID)
}

func (p Paypal) Refund(transactionID string) string {
	return fmt.Sprintf("Refunded transaction %s via PayPal", transactionID)
}

func (p Paypal) GetProviderName() string {
	return "PayPal"
}
