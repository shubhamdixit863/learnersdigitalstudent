package services

type PaymentProcessor interface {
	Pay(amount float64) (string, error)
	Refund(transactionID string) (string, error)
	GetProviderName() string
}
