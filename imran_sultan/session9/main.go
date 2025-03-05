package main

import (
	"fmt"
	"session9/internal/services"
)

type PayPal struct{ Email string }

type RazorPay struct{ UpiId string }

func ProssingMethod(p interface{}) {
	switch p.(type) {
	case PayPal:
		var service services.PaymentServices
		service = services.NewPayPal()
		fmt.Println(service.Pay(34))
		fmt.Println(service.Refund("ayz"))
		fmt.Println(service.GetProviderName())
	case RazorPay:
		var service services.PaymentServices
		service = services.NewRazorPay()
		fmt.Println(service.Pay(110))
		fmt.Println(service.Refund("xyz"))
		fmt.Println(service.GetProviderName())
	default:
		fmt.Println("Unknown payment processor")
	}
}

func main() {
	var choise int
	fmt.Println("Enter 1 for PayPal 2 for RazorPay")
	fmt.Scanln(&choise)
	switch choise {
	case 1:
		ProssingMethod(PayPal{"imran@paypal"})
	case 2:
		ProssingMethod(RazorPay{"imran@razorpay"})

	}
}
