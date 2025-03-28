package services

import (
	"fmt"
	"math/rand"
	"practice_payment/internals/storage"
	"practice_payment/internals/utils"
	"strconv"
)

type RazorPay struct {
	MerchId  string
	ApiKey   string
	ClientID string
}

func NewRazorPay() PaymentProcessor {

	return &RazorPay{
		MerchId:  "d222",
		ApiKey:   "3333",
		ClientID: "e333",
	}
}

func (razoepay *RazorPay) Pay(amount float64) (string, error) {

	if amount < 0 {
		return "", utils.CustomeErrorFuncRazorpay("amount is in negative")
	} else {
		trns := strconv.Itoa(rand.Int())
		obj := storage.NewTransaction(trns, "paypal")
		obj.AddTransaction(trns, "paypal")
		return "Paid using razor pay", nil

	}

}

func (razorpay *RazorPay) Refund(transactionId string) (string, error) {

	if transactionId == "" {
		return "", utils.CustomeErrorFuncRazorpay("transaction id invalid")
	} else {
		for _, x := range storage.List {
			if x.TransactionId == transactionId {
				fmt.Println(x)
				return "payment processing refund", nil

			}
		}
	}
	return "", utils.CustomeErrorFuncRazorpay("transactionid not found")
}

func (razor *RazorPay) GetProviderName() (string, error) {
	//fmt.Println("razor pay")

	return "razorpay", nil
}
