package utils

type CustomErrorPaypal struct {
	Message string
}

func (c *CustomErrorPaypal) Error() string {
	return c.Message + "by PayPal"
}

func CustomeErrorFuncPaypal(messsage string) error {
	return &CustomErrorPaypal{Message: messsage}

}

type CustomErrorStripe struct {
	Message string
}

func (c *CustomErrorStripe) Error() string {
	return c.Message + "by Stripe"
}
func CustomeErrorFuncStripe(messsage string) error {
	return &CustomErrorStripe{Message: messsage}

}

type CustomErrorRazorpay struct {
	Message string
}

func (c *CustomErrorRazorpay) Error() string {
	return c.Message + "by razorpay"
}
func CustomeErrorFuncRazorpay(messsage string) error {
	return &CustomErrorStripe{Message: messsage}
}
