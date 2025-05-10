package services

import (
	// "errors"
	"session8/practical/internal/utils"
)

type PaymentServices struct {
	Providers map[string]Services
}

// RegisterProvider(provider PaymentProcessor) → Adds a new payment provider.

// • ProcessPayment(providerName string, amount float64) (string, error) → Finds the provider and processes a payment.

// • IssueRefund(providerName, transactionID string) (string, error) → Processes a refund.

func NewPaymentServices() *PaymentServices {
	return &PaymentServices{
		Providers: make(map[string]Services),
	}
}

func (ps *PaymentServices) RegisterProvider(provider Services) {
	providerName := provider.GetProviderName()
	ps.Providers[providerName] = provider
}

func (ps *PaymentServices) ProcessPayment(providerName string, amount float64) (string, error) {

	
	provider, ok := ps.Providers[providerName]
   var err error
	if providerName=="Paypal"{
		err=&utils.PaypalErr{}
	}else if (providerName=="Stripe"){
		err=&utils.StripeErr{}
	}else if(providerName=="Razorpay"){
		err=&utils.RazorpayErr{}
	}
	if ok {
		return provider.Pay(amount), nil
	} else {
		return "",err
	}
}


func (ps *PaymentServices) IssueRefund(providerName string , transactionID string)(string,error){
	provider,ok := ps.Providers[providerName]
 var err error
	if providerName=="Paypal"{
		err=&utils.PaypalErr{}
	}else if (providerName=="Stripe"){
		err=&utils.StripeErr{}
	}else if(providerName=="Razorpay"){
		err=&utils.RazorpayErr{}
	}


	if ok {
		return provider.Refund(transactionID),nil
	}else {
		return "",err
	}
}

