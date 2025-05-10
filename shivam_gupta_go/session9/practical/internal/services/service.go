package services

// Pay(amount float64) string → Processes a payment.

// • Refund(transactionID string) string → Refunds a transaction.

// • GetProviderName() string → Returns the provider’s name.



type Services interface {
	Pay(amount float64) string
	Refund(transactionID string) string
	GetProviderName() string
}
