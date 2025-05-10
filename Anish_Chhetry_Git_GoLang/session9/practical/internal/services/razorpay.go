package services

import "fmt"

type Razorpay struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewRazorpay(client string, api string, merchant string) *Razorpay {
	return &Razorpay{
		ClientID:   client,
		APIKey:     api,
		MerchantID: merchant,
	}
}
func (s *Razorpay) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Razorpay (ClientID: %s, APIKey: %s, MerchantID: %s)", amount, s.ClientID, s.APIKey, s.MerchantID)
}
func (s *Razorpay) Refund(transactionID string) string {
	return fmt.Sprintf("Refunded %s via Razorpay", transactionID)
}
func (s *Razorpay) GetProviderName() string {
	return "Razorpay"
}
