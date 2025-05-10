package utils



type PaypalErr struct {
	
}

type StripeErr struct {
	
}
type RazorpayErr struct {
}

func(e *PaypalErr) Error() string{
	return "paypal : Not found"
}


func(e *StripeErr) Error() string{
	return " stripe : Not found"
}

func(e *RazorpayErr) Error() string{
	return "razorpay: Not found"
}
