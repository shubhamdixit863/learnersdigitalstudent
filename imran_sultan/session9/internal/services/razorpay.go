package services

import (
	"fmt"
	"session9/internal/util"
)

type RazorPay struct {
	UpiId string
}

func NewRazorPay() PaymentServices {
	fmt.Println("Welcome to Razorpay Gateway!...")
	fmt.Println("Enter your UpiId: ")
	var upiid string
	fmt.Scanln(&upiid)
	return &RazorPay{
		UpiId: upiid,
	}
}

func (p RazorPay) Pay(amount float64) (string, error) {
	reqamount := 99.9
	transid := "xyz"
	obj := NewTransiction()
	if amount < reqamount {
		return "", &util.PaypalError{Message: "Transiction failed not enough money!"}
	}
	obj.AddNewTransiction(p.UpiId, amount, transid, "Razorpay")
	return fmt.Sprintln("You paid amount", amount, " RazorPay id: ", p.UpiId), nil
}
func (p RazorPay) Refund(transactionID string) (string, error) {
	for i := 0; i < len(Mydata); i++ {
		if Mydata[i].Transid == transactionID {
			Mydata[i].Amount = 0
			return "Refunded successfully!", nil
		}
	}

	return "", &util.RazorPayError{Message: "transiction Not found!"}
}
func (p RazorPay) GetProviderName() string {
	return "Payment done via RazorPay!"
}
