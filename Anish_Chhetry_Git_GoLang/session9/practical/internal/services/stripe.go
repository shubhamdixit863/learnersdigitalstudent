package services

import "fmt"

type Stripe struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewStripe(client string, api string, merchant string) *PayPal {
	return &PayPal{
		ClientID:   client,
		APIKey:     api,
		MerchantID: merchant,
	}
}
func (s *Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Stripe (ClientID: %s, APIKey: %s, MerchantID: %s)", amount, s.ClientID, s.APIKey, s.MerchantID)
}
func (s *Stripe) Refund(transactionID string) string {
	return fmt.Sprintf("Refunded %s via Stripe", transactionID)
}
func (s *Stripe) GetProviderName() string {
	return "Stripe"
}
