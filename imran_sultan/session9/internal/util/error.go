package util

type PaypalError struct {
	Message string
}

func (p *PaypalError) Error() string {
	return p.Message
}

type RazorPayError struct {
	Message string
}

func (r *RazorPayError) Error() string {
	return r.Message
}
