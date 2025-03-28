package services

type PaymentProcessor interface {
	Pay(amount float64) (string, error)
	Refund(transactionId string) (string, error)
	GetProviderName() (string, error)
}

func ProviderDetails(provider PaymentProcessor) {
	//switch provider.(type) {
	//case *Paypal:
	//	fmt.Println(provider.GetProviderName())
	//
	//case *Stripe:
	//	fmt.Println(provider.GetProviderName())
	//case *RazorPay:
	//	fmt.Println(provider.GetProviderName())
	//
	//}
}
