package services

import "fmt"

type PayPalService struct {
	APIKey string
}

func NewPayPalService(apikey string) PaymentProcessor {
	return &PayPalService{
		APIKey: apikey,
	}
}

func (p *PayPalService) Pay(amount float64) string {
	if amount < 0 {
		return ""
	}
	return fmt.Sprintf("PayPal Payment of %.2f processed.", amount)
}
func (p *PayPalService) Refund(transactionID string) string {
	return fmt.Sprintf("PayPal Refund for TranscationID %s processed.", transactionID)
}
func (p *PayPalService) GetProviderName() string {
	return "PayPal"
}
func (p *PayPalService) GetAPIKey() string {
	return p.APIKey
}
