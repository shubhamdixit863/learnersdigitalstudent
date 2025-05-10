package services

import (
	"fmt"
	"practical/internal/utils"
)

type PaymentGateway struct {
	paymentProviders []PaymentProcessor
}

func (r *PaymentGateway) RegisterProvider(provider PaymentProcessor) {
	r.paymentProviders = append(r.paymentProviders, provider)
	fmt.Println("Registered", provider)
}

func (r *PaymentGateway) ProcessPayment(providerName string, amount float64) (string, error) {
	for _, v := range r.paymentProviders {
		if v.GetProviderName() == providerName {
			return v.Pay(amount)
		}
	}
	return "", utils.NewProviderNotFoundError()
}
func (r *PaymentGateway) IssueRefund(providerName, transactionID string) (string, error) {
	for _, v := range r.paymentProviders {
		if v.GetProviderName() == providerName {
			return v.Refund(transactionID)
		}
	}
	return "", utils.NewProviderNotFoundError()
}

func ExtractProvider(providerName PaymentProcessor) (string, error) {
	switch p := providerName.(type) {
	case *PayPal:
		return fmt.Sprintf("PayPal: ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID), nil
	case *Stripe:
		return fmt.Sprintf("Stripe: ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID), nil
	case *Razorpay:
		return fmt.Sprintf("Razorpay: ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID), nil
	default:
		return "", utils.NewProviderNotFoundError()
	}
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{
		paymentProviders: []PaymentProcessor{},
	}
}
