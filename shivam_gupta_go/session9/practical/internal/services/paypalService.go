package services

import "fmt"

// Create three struct implementations: PayPal, Stripe, and Razorpay.

// • Each struct should:

// • Store unique details (e.g., ClientID, APIKey, MerchantID).

// • Implement the PaymentProcessor interface.

type PayPal struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal (ClientID: %s)", amount, p.ClientID)
}

func (p *PayPal) Refund(transactionID string) string {
	return fmt.Sprintf("Refund issued for transaction %s via PayPal", transactionID)
}

func (p *PayPal) GetProviderName() string {
	return "PayPal"
}


