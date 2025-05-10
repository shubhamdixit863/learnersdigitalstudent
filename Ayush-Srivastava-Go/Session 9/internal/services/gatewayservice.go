package services

import (
	"payment-gateway/internal/util"
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
	provider, ok := pg.providers[providerName]
	if !ok {

		paymentError := util.NewPaymentError("Error while payment", providerName)

		return "", paymentError
	}
	return provider.Pay(amount), nil
}

func (pg *PaymentGateway) IssueRefund(providerName string, transactionID string) (string, error) {
	provider, ok := pg.providers[providerName]
	if !ok {

		paymentError := util.NewPaymentError("Error while refund", providerName)

		return "", paymentError
	}
	return provider.Refund(transactionID), nil
}
