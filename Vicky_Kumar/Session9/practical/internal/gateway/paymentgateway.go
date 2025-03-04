package gateway

import (
	"fmt"
	"practical/internal/services"
	"practical/internal/transactions"
	"practical/internal/utils"
)

type PaymentGateway struct {
	providers map[string]services.PaymentProcessor
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{
		providers: make(map[string]services.PaymentProcessor),
	}
}
func (pg *PaymentGateway) RegisterProvider(provider services.PaymentProcessor) {
	pg.providers[provider.GetProviderName()] = provider
}

func (pg *PaymentGateway) ProcessPayment(providerName string, amount float64) (string, error) {
	CID := utils.RandomID()
	MID := utils.RandomID()
	TID := utils.RandomID()
	var transaction = transactions.Transaction{
		TranscationID:   TID,
		ClientID:        CID,
		MerchantID:      MID,
		PaymentProvider: providerName,
	}
	transactions.Transactions = append(transactions.Transactions, transaction)
	detail := pg.providers[providerName].Pay(amount)
	if detail == "" {
		return "", fmt.Errorf("Some Error Occured. Try Again!")
	}
	return detail, nil
}

func (pg *PaymentGateway) IssueRefund(providerName, transactionID string) (string, error) {
	detail := pg.providers[providerName].Refund(transactionID)
	return detail, nil
}
func (pg *PaymentGateway) ProviderDetail(provider services.PaymentProcessor) {
	fmt.Println("Provider Details:")
	switch provider.(type) {
	case *services.PayPalService:
		fmt.Println(provider.GetProviderName())

	case *services.RazorPayService:
		fmt.Println(provider.GetProviderName())

	case *services.StripeService:
		fmt.Println(provider.GetProviderName())
	}
	fmt.Printf("API Key: %s\n", provider.GetAPIKey())
}
