package payment

import (
	"errors"
	"fmt"
)

type PaymentGateway struct {
	providers map[string]PaymentProcessor
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{providers: make(map[string]PaymentProcessor)}
}

func (pg *PaymentGateway) RegisterProvider(provider PaymentProcessor) {
	pg.providers[provider.GetProviderName()] = provider
}

func (pg *PaymentGateway) ProcessPayment(providerName string, amount float64) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", errors.New("payment provider not registered")
	}
	return provider.Pay(amount), nil
}

func (pg *PaymentGateway) IssueRefund(providerName, transactionID string) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", errors.New("payment provider not registered")
	}
	return provider.Refund(transactionID), nil
}

func ExtractProviderDetails(provider PaymentProcessor) {
	fmt.Println("Extracting provider details...")

	switch p := provider.(type) {
	case PayPal:
		fmt.Printf("Provider: PayPal, ClientID: %s, APIKey: %s, \n", p.ClientID, p.APIKey)
	case Stripe:
		fmt.Printf("Provider: Stripe, APIKey: %s\n", p.APIKey)
	case Razorpay:
		fmt.Printf("Provider: Razorpay, MerchantID: %s, APIKey: %s\n", p.MerchantID, p.APIKey)
	default:
		fmt.Println("Unknown provider")
	}
}
