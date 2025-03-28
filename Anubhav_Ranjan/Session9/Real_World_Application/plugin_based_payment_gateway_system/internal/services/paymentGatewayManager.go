package services

import (
	"fmt"
	"plugin_based_payment_gateway_system/internal/customErrors"
	"plugin_based_payment_gateway_system/internal/paymentInterface"
	"plugin_based_payment_gateway_system/internal/paymentProviders"
	"plugin_based_payment_gateway_system/internal/transactions"
	"reflect"

	"github.com/google/uuid"
)

type PaymentGatewayManager struct {
	providers    map[string]paymentInterface.PaymentProcessor
	transactions map[string]transactions.Transaction
}

func NewPaymentGateway() *PaymentGatewayManager {
	return &PaymentGatewayManager{
		providers:    make(map[string]paymentInterface.PaymentProcessor),
		transactions: make(map[string]transactions.Transaction),
	}
}

func (pg *PaymentGatewayManager) RegisterProvider(provider paymentInterface.PaymentProcessor) {
	pg.providers[provider.GetProviderName()] = provider
	fmt.Println(provider.GetProviderName(), "registered successfully.")
}

func (pg *PaymentGatewayManager) ProcessPayment(providerName string, amount float64) (string, error) {
	provider, exists := pg.providers[providerName]
	if !exists {
		return "", customErrors.NewGatewayError("payment provider not found", nil)
	}

	if amount < 0 {
		return "", customErrors.NewGatewayError("Amount Must be Greater than Zero..", provider)
	}

	transactionID := uuid.New().String()[:8]
	status := provider.Pay(amount)

	transaction := transactions.Transaction{
		TransactionID: transactionID,
		Amount:        amount,
		ProviderName:  providerName,
		Status:        "Completed",
	}

	pg.transactions[transactionID] = transaction
	fmt.Println("Payment Successful:", transaction.String())
	return status, nil
}

func (pg *PaymentGatewayManager) IssueRefund(transactionID string) (string, error) {
	transaction, exists := pg.transactions[transactionID]
	if !exists {
		return "", customErrors.NewGatewayError("transaction not found", nil)
	}

	provider, exists := pg.providers[transaction.ProviderName]
	if !exists {
		return "", customErrors.NewGatewayError("payment provider not found", nil)
	}

	refundStatus := provider.Refund(transactionID)
	transaction.Status = "Refunded"
	pg.transactions[transactionID] = transaction

	fmt.Println("Refund Successful:", transaction.String())
	return refundStatus, nil
}

func (pg *PaymentGatewayManager) DisplayProviderDetails(providerName string) {
	provider, exists := pg.providers[providerName]
	if !exists {
		fmt.Println("Provider not found")
		return
	}

	fmt.Println(reflect.TypeOf(provider))

	switch p := provider.(type) {
	case *paymentProviders.PayPal:
		fmt.Println("PayPal:", p)
	case *paymentProviders.Stripe:
		fmt.Println("Stripe:", p)
	case *paymentProviders.RazorPay:
		fmt.Println("Razorpay:", p)
	default:
		fmt.Println("Unknown provider")
	}
}
