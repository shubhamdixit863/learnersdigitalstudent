package paymentProviders

import (
	"fmt"
	"plugin_based_payment_gateway_system/internal/utils"
)

type RazorPay struct {
	ClientID   string
	APIKey     string
	MerchantID string
}

func NewRazorPay() *RazorPay {
	return &RazorPay{
		ClientID:   utils.GenerateID("RazorPay-Client"),
		APIKey:     utils.GenerateID("RazorPay-API"),
		MerchantID: utils.GenerateID("RazorPay-Merchant"),
	}
}

func (p *RazorPay) Pay(amount float64) string {
	return fmt.Sprintf("Razorpay: Processed payment of $%.2f", amount)
}

func (p *RazorPay) Refund(transactionID string) string {
	return fmt.Sprintf("Razorpay: Refunded transaction %s", transactionID)
}

func (p *RazorPay) GetProviderName() string {
	return "RazorPay"
}

func (r *RazorPay) String() string {
	return fmt.Sprintf("Client ID: %s, APIKey: %s, MerchantID: %s",
		r.ClientID, r.APIKey, r.MerchantID)
}
