package paymentProviders

import (
	"fmt"
	"plugin_based_payment_gateway_system/internal/utils"
)

type Stripe struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewStripe() *Stripe {
	return &Stripe{
		ClientID:   utils.GenerateID("Stripe-Client"),
		APIKey:     utils.GenerateID("Stripe-API"),
		MerchantID: utils.GenerateID("Stripe-Merchant"),
	}
}

func (p *Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Stripe: Processed payment of $%.2f", amount)
}

func (p *Stripe) Refund(transactionID string) string {
	return fmt.Sprintf("Stripe: Refunded transaction %s", transactionID)
}

func (p *Stripe) GetProviderName() string {
	return "Stripe"
}

func (s *Stripe) String() string {
	return fmt.Sprintf("Client ID: %s, APIKey: %s, MerchantID: %s",
		s.ClientID, s.APIKey, s.MerchantID)
}
