package services

type PaymentServices interface {
	Pay(float64) (string, error)
	Refund(string) (string, error)
	GetProviderName() string
}
