package gateway

import (
	"errors"
	"handsOn/internal/services"
)

// PaymentGateway manages multiple payment providers
type PaymentGateway struct {
	providers map[string]services.PaymentProcessor
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{providers: make(map[string]services.PaymentProcessor)}
}

func (pg *PaymentGateway) RegisterProvider(provider services.PaymentProcessor) {
	pg.providers[provider.GetProviderName()] = provider
}

func (pg *PaymentGateway) ProcessPayment(providerName string, amount float64) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", errors.New("payment provider not found")
	}
	transaction, err := provider.Pay(amount)
	return transaction, err
}

func (pg *PaymentGateway) IssueRefund(providerName, transactionID string) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", errors.New("payment provider not found")
	}
	transaction, err := provider.Refund(transactionID)
	return transaction, err
}
