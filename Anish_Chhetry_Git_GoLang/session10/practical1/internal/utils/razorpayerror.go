package utils

// AMOUNT ERROR
type RazorpayAmountError struct {
}

func (c *RazorpayAmountError) Error() string {
	return "RAZORPAY ERROR: Please Enter Amount more than 0"
}
func NewRazorpayAmountZeroError() error {
	return &RazorpayAmountError{}
}

// TRANSACTION ID ERROR
type RazorpayTransactionIDError struct {
}

func (c *RazorpayTransactionIDError) Error() string {
	return "RAZORPAY ERROR: Please Enter the transaction ID"
}
func NewRazorpayEmptyTransactionIDError() error {
	return &RazorpayTransactionIDError{}
}
