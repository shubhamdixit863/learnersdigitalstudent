package services

import "fmt"

type StripeError struct {
	Operation string
	Reason    string
}

func NewStripeError(operation, reason string) *StripeError {
	return &StripeError{Operation: operation, Reason: reason}
}

func (e *StripeError) Error() string {
	return fmt.Sprintf("Razorpay Error - Operation: %s, Reason: %s", e.Operation, e.Reason)
}

type Stripe struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewStripe(clientID, apiKey, merchantID string) PaymentProcessor {
	return &Stripe{
		ClientID:   clientID,
		APIKey:     apiKey,
		MerchantID: merchantID,
	}
}

func (p Stripe) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", NewStripeError("Pay", "Invalid amount. Payment must be greater than zero.")
	}
	return fmt.Sprintf("Paid $%.2f via  Stripe (ClientID: %s)", amount, p.ClientID), nil
}

func (p Stripe) Refund(transactionID string) (string, error) {
	if transactionID == "" {
		return "", NewPaypalError("Refund", "Transaction ID cannot be empty.")
	}
	return fmt.Sprintf("Refunded transaction %s via Stripe", transactionID), nil
}

func (p Stripe) GetProviderName() string {
	return "Stripe"
}
