package services

import (
	"fmt"
	"payment_gateway/internal/utils"
)

type PaymentGateway struct {
	providers map[string]PaymentProcessor
	Error utils.Payment_Error
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{providers: make(map[string]PaymentProcessor), Error: utils.Payment_Error{}}
}

func (pg *PaymentGateway) RegisterProvider(provider PaymentProcessor) {
	pg.providers[provider.GetProviderName()] = provider
}

func (pg *PaymentGateway) ProcessPayment(providerName string, amount float64) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", &utils.Payment_Error{
			Msg:    "payment provider not registered",
			P_name: providerName,
		}
	}
	return provider.Pay(amount), nil
}

func (pg *PaymentGateway) IssueRefund(providerName, transactionID string) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", &utils.Payment_Error{
			Msg:    "payment provider not registered",
			P_name: providerName,
		}
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