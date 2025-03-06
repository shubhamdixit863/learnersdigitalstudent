package services

import (
	"fmt"
	"practical/internal/utils"
)

type Stripe struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewStripe(client string, api string, merchant string) *Stripe {
	return &Stripe{
		ClientID:   client,
		APIKey:     api,
		MerchantID: merchant,
	}
}
func (s *Stripe) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", utils.NewStripeAmountZeroError()
	}
	return fmt.Sprintf("Paid %.2f using Stripe (ClientID: %s, APIKey: %s, MerchantID: %s)", amount, s.ClientID, s.APIKey, s.MerchantID), nil
}
func (s *Stripe) Refund(transactionID string) (string, error) {
	if transactionID == "" {
		return "", utils.NewStripeEmptyTransactionIDError()
	}
	return fmt.Sprintf("Refunded %s via Stripe", transactionID), nil
}
func (s *Stripe) GetProviderName() string {
	return "Stripe"
}
