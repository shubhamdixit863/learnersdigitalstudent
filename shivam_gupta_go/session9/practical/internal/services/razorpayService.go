package services

import "fmt"

type Razorpay struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func (p *Razorpay) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Razorpay (clientid : %s )",amount,p.ClientID)
}

func (p *Razorpay) Refund(transactionID string )string {
	return fmt.Sprintf("Refund issued for transaction %s via Rozarpay",transactionID)
}

func (p *Razorpay) GetProviderName()string{
	return "Rozarpay"
}