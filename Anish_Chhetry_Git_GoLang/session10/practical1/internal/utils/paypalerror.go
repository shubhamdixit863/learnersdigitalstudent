package utils

// AMOUNT ERROR
type PaypalAmountError struct {
}

func (c *PaypalAmountError) Error() string {
	return "PAYPAL ERROR: Please Enter Amount more than 0"
}
func NewPaypalAmountZeroError() error {
	return &PaypalAmountError{}
}

// TRANSACTION ID ERROR
type PaypalTransactionIDError struct {
}

func (c *PaypalTransactionIDError) Error() string {
	return "PAYPAL ERROR: Please Enter the transaction ID"
}
func NewPaypalEmptyTransactionIDError() error {
	return &PaypalTransactionIDError{}
}
