package services

import "fmt"

type PayPal struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewPaypal(client string, api string, merchant string) *PayPal {
	return &PayPal{
		ClientID:   client,
		APIKey:     api,
		MerchantID: merchant,
	}
}
func (s *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal (ClientID: %s, APIKey: %s, MerchantID: %s)", amount, s.ClientID, s.APIKey, s.MerchantID)
}
func (s *PayPal) Refund(transactionID string) string {
	return fmt.Sprintf("Refunded %s via PayPal", transactionID)
}
func (s *PayPal) GetProviderName() string {
	return "PayPal"
}
