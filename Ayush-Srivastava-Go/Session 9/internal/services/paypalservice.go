package services

import "strconv"

type PayPal struct {
	ClientID string
	APIKey   string
	MerchantID string
}

func (p PayPal) Pay(amount float64) string {
	return "Paid an amount of " + strconv.FormatFloat(amount, 'f', 2, 64) + " using PayPal"
}

func (p PayPal) Refund(transactionID string) string {
	return "Refunded transaction ID " + transactionID + " using PayPal"
}

func (p PayPal) GetProviderName() string {
	return "PayPal"
}
