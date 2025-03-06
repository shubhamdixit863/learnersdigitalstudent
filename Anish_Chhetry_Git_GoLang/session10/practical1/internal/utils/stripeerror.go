package utils

// AMOUNT ERROR
type StripeAmountError struct {
}

func (c *StripeAmountError) Error() string {
	return "STRIPE ERROR: Please Enter Amount more than 0"
}
func NewStripeAmountZeroError() error {
	return &StripeAmountError{}
}

// TRANSACTION ID ERROR
type StripeTransactionIDError struct {
}

func (c *StripeTransactionIDError) Error() string {
	return "STRIPE ERROR: Please Enter the transaction ID"
}
func NewStripeEmptyTransactionIDError() error {
	return &StripeTransactionIDError{}
}
