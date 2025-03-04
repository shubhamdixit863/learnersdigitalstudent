package services

type PaymentProcessor interface {
	Pay(amount float64) string
	Refund(transactionID string) string
	GetProviderName() string
	GetAPIKey() string
}
