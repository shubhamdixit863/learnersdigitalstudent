package services

import "fmt"

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
			return v.Pay(amount), nil
		}
	}
	return "Provider Not Found", nil
}
func (r *PaymentGateway) IssueRefund(providerName, transactionID string) (string, error) {
	for _, v := range r.paymentProviders {
		if v.GetProviderName() == providerName {
			return v.Refund(transactionID), nil
		}
	}
	return "Provider Not Found", nil
}

func ExtractProvider(providerName PaymentProcessor) {
	switch p := providerName.(type) {
	case *PayPal:
		fmt.Printf("PayPal: ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	case *Stripe:
		fmt.Printf("Stripe: ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	case *Razorpay:
		fmt.Printf("Razorpay:  ClientID: %s, APIKey: %s, MerchantID: %s\n", p.ClientID, p.APIKey, p.MerchantID)
	default:
		fmt.Println("Unknown provider")
	}
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{
		paymentProviders: []PaymentProcessor{},
	}
}
