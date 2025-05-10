package services

import "strconv"

type Stripe struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func (s Stripe) Pay(amount float64) string {
	return "Paid an amount of " + strconv.FormatFloat(amount, 'f', 2, 64) + " using Stripe"
}

func (s Stripe) Refund(transactionID string) string {
	return "Refunded transaction ID " + transactionID + " using Stripe"
}

func (s Stripe) GetProviderName() string {
	return "Stripe"
}