package paymentProviders

import (
	"fmt"
	"plugin_based_payment_gateway_system/internal/utils"
)

type PayPal struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewPayPal() *PayPal {
	return &PayPal{
		ClientID:   utils.GenerateID("PayPal-Client"),
		APIKey:     utils.GenerateID("PayPal-API"),
		MerchantID: utils.GenerateID("PayPal-Merchant"),
	}
}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("PayPal: Processed payment of $%.2f", amount)
}

func (p *PayPal) Refund(transactionID string) string {
	return fmt.Sprintf("PayPal: Refunded transaction %s", transactionID)
}

func (p *PayPal) GetProviderName() string {
	return "PayPal"
}

func (p *PayPal) String() string {
	return fmt.Sprintf("Client ID: %s, APIKey: %s, MerchantID: %s",
		p.ClientID, p.APIKey, p.MerchantID)
}
