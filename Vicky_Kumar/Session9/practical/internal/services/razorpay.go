package services

import "fmt"

type RazorPayService struct {
	APIKey string
}

func NewRazorPayService(apikey string) PaymentProcessor {
	return &RazorPayService{
		APIKey: apikey,
	}
}

func (r *RazorPayService) Pay(amount float64) string {
	if amount < 0 {
		return ""
	}
	return fmt.Sprintf("RazorPay Payment of %.2f processed.", amount)
}
func (r *RazorPayService) Refund(transactionID string) string {
	return fmt.Sprintf("RazorPay Refund for TranscationID %s processed.", transactionID)
}
func (r *RazorPayService) GetProviderName() string {
	return "RazorPay"
}
func (r *RazorPayService) GetAPIKey() string {
	return r.APIKey
}
