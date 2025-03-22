package services

import "strconv"

type RazorPay struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func (r RazorPay) Pay(amount float64) string {
	return "Paid an amount of " + strconv.FormatFloat(amount, 'f', 2, 64) + " using RazorPay"
}

func (r RazorPay) Refund(transactionID string) string {
	return "Refunded transaction ID " + transactionID + " using RazorPay"
}

func (r RazorPay) GetProviderName() string {
	return "RazorPay"
}