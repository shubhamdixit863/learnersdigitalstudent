package services

import "fmt"

type PaypalError struct {
	Operation string
	Reason    string
}

func NewPaypalError(operation, reason string) *PaypalError {
	return &PaypalError{Operation: operation, Reason: reason}
}

func (e *PaypalError) Error() string {
	return fmt.Sprintf("PayPal Error - Operation: %s, Reason: %s", e.Operation, e.Reason)
}

type Paypal struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewPaypal(clientID, apiKey, merchantID string) PaymentProcessor {
	return &Paypal{
		ClientID:   clientID,
		APIKey:     apiKey,
		MerchantID: merchantID,
	}
}
func (p Paypal) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", NewPaypalError("Pay", "Invalid amount. Payment must be greater than zero.")
	}
	return fmt.Sprintf("Paid $%.2f via PayPal (ClientID: %s)", amount, p.ClientID), nil
}

func (p Paypal) Refund(transactionID string) (string, error) {
	if transactionID == "" {
		return "", NewPaypalError("Refund", "Transaction ID cannot be empty.")
	}
	return fmt.Sprintf("Refunded transaction %s via PayPal", transactionID), nil
}

func (p Paypal) GetProviderName() string {
	return "PayPal"
}
