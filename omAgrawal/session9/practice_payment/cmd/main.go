package main

import (
	"fmt"
	"practice_payment/internals/services"
)

func main() {
	var paymentProvider services.PaymentProcessor

	//pg := services.NewPaymentGateway()

	for {
		fmt.Println("select a 1.Pay\n2.Refund\n3getProviderDetails")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("1.Paypal\n2.Stripe\n3.razorPay")
			var provider int
			fmt.Scanln(&provider)

			var amount float64
			fmt.Println("enter amount")
			fmt.Scanln(&amount)

			switch provider {

			case 1:
				paymentProvider = services.NewPaypal("jnjwj", "njfewjnnjw", "dwde")
				c, err := paymentProvider.Pay(amount)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(c)
			case 2:
				paymentProvider = services.NewStripe()
				c, err := paymentProvider.Pay(amount)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(c)
				}
			case 3:
				paymentProvider = services.NewRazorPay()
				c, err := paymentProvider.Pay(amount)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(c)
				}
			}

		case 2:
			fmt.Println("Enter transaction id")
			var transactionId string
			fmt.Scanln(&transactionId)
			paymentProvider.Refund(transactionId)
		case 3:
			//var provier string
			//fmt.Println("enter the payment provider name")
			//fmt.Scanln(&provier)
			paymentProvider = services.NewStripe()
			services.ProviderDetails(paymentProvider)

		}

	}

}
