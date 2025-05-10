// package main

// import (
// 	"errors"
// 	"log"
// )

// type Error interface {
// 	Error() string
// }

// type CustomError struct {
// 	message string
// }

// func (c *CustomError) Error() string {
// 	return c.message
// }

// func Divide(a, b int) (int, error) {
// 	if b == 0 {
// 		err := errors.New("Cannot Divide by 0")
// 		return 0, &CustomError{message: err.Error()}
// 	}
// 	return a / b, nil
// }

// func ReturnCustomError() error {
// 	return &CustomError{message: "This is a custom error"}
// }

// func main() {
// 	b, err := Divide(6, 0)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(b)
// }


package main

import (
	"fmt"
	"log"
	"session10/internal/services"
)

func main() {
	gateway := services.NewPaymentGateway()

	paypal := &services.PayPal{ClientID: "paypal123", APIKey: "123API", MerchantID: "SGSDSJD"}
	gateway.RegisterProvider(paypal)

	stripe:=&services.Stripe{ClientID: "stripe123", APIKey: "123A", MerchantID: "SGSDS"}
	gateway.RegisterProvider(stripe)

	paymentMsg, err := gateway.ProcessPayment("PayPal", -50.00)
	if err != nil {
		log.Println("Payment failed:", err)
	} else {
		fmt.Println(paymentMsg)
	}

	refundMsg, err := gateway.IssueRefund("PayPal", "")
	if err != nil {
		log.Println("Refund failed:", err)
	} else {
		fmt.Println(refundMsg)
	}

	provider := gateway.Providers["PayPal"]
	services.GetProviderDetails(provider)


	paymentM,err := gateway.ProcessPayment("Stripe",500)
	if err!=nil{
		log.Println("Payment Fail",err)
	}else{
		fmt.Println(paymentM)
	}
	refundM,err:=gateway.IssueRefund("Stripe","")
	if err!=nil{
		log.Println("Payment Fail",err)
	}else{
		fmt.Println(refundM)
	}

	services.GetProviderDetails(stripe)




}

