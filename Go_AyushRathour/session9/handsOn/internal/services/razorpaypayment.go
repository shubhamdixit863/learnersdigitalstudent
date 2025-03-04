package services

import "fmt"

type Razorpay struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewRazorpay(clientID, apiKey, merchantID string) PaymentProcessor {
	return &Razorpay{
		ClientID:   clientID,
		APIKey:     apiKey,
		MerchantID: merchantID,
	}
}

func (p Razorpay) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f via RazorPay (ClientID: %s)", amount, p.ClientID)
}

func (p Razorpay) Refund(transactionID string) string {
	return fmt.Sprintf("Refunded transaction %s via RazorPay", transactionID)
}

func (p Razorpay) GetProviderName() string {
	return "Razorpay"
}
