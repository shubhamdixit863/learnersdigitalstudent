package services

import (
	"fmt"
	"practical/internal/utils"
)

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
func (s *PayPal) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", utils.NewPaypalAmountZeroError()
	}
	return fmt.Sprintf("Paid %.2f using PayPal (ClientID: %s, APIKey: %s, MerchantID: %s)", amount, s.ClientID, s.APIKey, s.MerchantID), nil
}
func (s *PayPal) Refund(transactionID string) (string, error) {
	if transactionID == "" {
		return "", utils.NewPaypalEmptyTransactionIDError()
	}
	return fmt.Sprintf("Refunded %s via PayPal", transactionID), nil
}
func (s *PayPal) GetProviderName() string {
	return "PayPal"
}
