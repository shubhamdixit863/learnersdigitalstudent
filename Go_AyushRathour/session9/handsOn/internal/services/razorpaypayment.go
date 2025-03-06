package services

import "fmt"

type RazorpayError struct {
	Operation string
	Reason    string
}

func NewRazorpayError(operation, reason string) *RazorpayError {
	return &RazorpayError{Operation: operation, Reason: reason}
}

func (e *RazorpayError) Error() string {
	return fmt.Sprintf("Razorpay Error - Operation: %s, Reason: %s", e.Operation, e.Reason)
}

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

func (p Razorpay) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", NewPaypalError("Pay", "Invalid amount. Payment must be greater than zero.")
	}
	return fmt.Sprintf("Paid $%.2f via PayPal (ClientID: %s)", amount, p.ClientID), nil
}

func (p Razorpay) Refund(transactionID string) (string, error) {
	if transactionID == "" {
		return "", NewPaypalError("Refund", "Transaction ID cannot be empty.")
	}
	return fmt.Sprintf("Refunded transaction %s via RazorPay", transactionID), nil
}

func (p Razorpay) GetProviderName() string {
	return "Razorpay"
}
