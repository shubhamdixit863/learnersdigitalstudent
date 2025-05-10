package services

import "fmt"

type StripeService struct {
	APIKey string
}

func NewStripeService(apikey string) PaymentProcessor {
	return &StripeService{
		APIKey: apikey,
	}
}

func (s *StripeService) Pay(amount float64) string {
	if amount < 0 {
		return ""
	}
	return fmt.Sprintf("Stripe Payment of %.2f processed.", amount)
}
func (s *StripeService) Refund(transactionID string) string {
	return fmt.Sprintf("Stripe Refund for TranscationID %s processed.", transactionID)
}
func (s *StripeService) GetProviderName() string {
	return "Stripe"
}
func (s *StripeService) GetAPIKey() string {
	return s.APIKey
}
