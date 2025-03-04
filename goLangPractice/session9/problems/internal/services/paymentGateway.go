package services

import (
	"errors"
	"fmt"
)

type PaymentGateway struct {
	providers map[string]PaymentProcessor
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{
		providers: make(map[string]PaymentProcessor),
	}
}

func (pg *PaymentGateway) RegisterProvider(provider PaymentProcessor) {
	pg.providers[provider.GetProviderName()] = provider
}

func (pg *PaymentGateway) ProcessPayment(providerName string, amount float64) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", errors.New("provider not registered")
	}
	return provider.Pay(amount), nil
}

func (pg *PaymentGateway) IssueRefund(providerName, transactionID string) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", errors.New("provider not registered")
	}
	return provider.Refund(transactionID), nil
}

func ExtractProviderDetails(provider PaymentProcessor) {
	switch p := provider.(type) {
	case *PayPal:
		fmt.Printf("PayPal ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	case *Stripe:
		fmt.Printf("Stripe ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	case *Razorpay:
		fmt.Printf("Razorpay ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	default:
		fmt.Println("Unknown provider")
	}
}
