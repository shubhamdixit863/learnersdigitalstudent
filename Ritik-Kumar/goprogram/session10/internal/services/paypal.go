package services

import (
	"fmt"
	"session10/internal/utils"
)

type PayPal struct { 
	ClientID   string
	APIKey     string
	MerchantID string
}

func (p *PayPal) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", utils.NewPaymentError("PayPal", "Invalid payment amount")
	}
	return fmt.Sprintf("PayPal processed payment of $%.2f", amount), nil
}

func (p *PayPal) Refund(transactionID string) (string, error) {
	if transactionID == "" {
		return "", utils.NewPaymentError("PayPal", "Invalid transaction ID for refund")
	}
	return fmt.Sprintf("PayPal refunded transaction %s", transactionID), nil
}

func (p *PayPal) GetProviderName() string {
	return "PayPal"
}


type Stripe struct { 
	ClientID   string
	APIKey     string
	MerchantID string
}

func (p *Stripe) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", utils.NewStripeError("Stripe", "Invalid payment amount")
	}
	return fmt.Sprintf("Stripe processed payment of $%.2f", amount), nil
}

func (p *Stripe) Refund(transactionID string) (string, error) {
	if transactionID == "" {
		return "", utils.NewStripeError("Stripe", "Invalid transaction ID for refund")
	}
	return fmt.Sprintf("Stripe refunded transaction %s", transactionID), nil
}

func (p *Stripe) GetProviderName() string {
	return "Stripe"
}
