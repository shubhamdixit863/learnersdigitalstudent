package services

import (
	"fmt"
	"practical/internal/utils"
)

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
func (s *Razorpay) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", utils.NewRazorpayAmountZeroError()
	}
	return fmt.Sprintf("Paid %.2f using Razorpay (ClientID: %s, APIKey: %s, MerchantID: %s)", amount, s.ClientID, s.APIKey, s.MerchantID), nil
}
func (s *Razorpay) Refund(transactionID string) (string, error) {
	if transactionID == "" {
		return "", utils.NewRazorpayEmptyTransactionIDError()
	}
	return fmt.Sprintf("Refunded %s via Razorpay", transactionID), nil
}
func (s *Razorpay) GetProviderName() string {
	return "Razorpay"
}
