package services

import "errors"

type PaymentGateway struct {
	Providers map[string]PaymentProcessor
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{Providers: make(map[string]PaymentProcessor)}
}

func (pg *PaymentGateway) RegisterProvider(provider PaymentProcessor) {
	pg.Providers[provider.GetProviderName()] = provider
}

func (pg *PaymentGateway) ProcessPayment(providerName string, amount float64) (string, error) {
	provider, exists := pg.Providers[providerName]
	if !exists {
		return "", errors.New("provider not found")
	}
	return provider.Pay(amount), nil
}

func (pg *PaymentGateway) IssueRefund(providerName, transactionID string) (string, error) {
	provider, exists := pg.Providers[providerName]
	if !exists {
		return "", errors.New("provider not found")
	}
	return provider.Refund(transactionID), nil
}
