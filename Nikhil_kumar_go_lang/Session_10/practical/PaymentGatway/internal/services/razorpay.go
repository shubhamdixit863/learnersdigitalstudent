package services

import "fmt"

type Razorpay struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func (r Razorpay) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Razorpay", amount)
}

func (r Razorpay) Refund(transactionID string) string {
	return fmt.Sprintf("Refund issued for transaction %s via Razorpay", transactionID)
}

func (r Razorpay) GetProviderName() string {
	return "Razorpay"
}
