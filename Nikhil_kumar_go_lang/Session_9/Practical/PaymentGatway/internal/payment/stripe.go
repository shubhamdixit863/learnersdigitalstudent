package payment

import "fmt"

type Stripe struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func (s Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Stripe", amount)
}
func (s Stripe) Refund(transactionID string) string {
	return fmt.Sprintf("Refunded transaction %s using Stripe", transactionID)
}
func (s Stripe) GetProviderName() string {
	return "Stripe"
}
