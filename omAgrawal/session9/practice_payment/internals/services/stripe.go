package services

import (
	"fmt"
	"math/rand"
	"practice_payment/internals/storage"
	"practice_payment/internals/utils"
	"strconv"
)

type Stripe struct {
	ClientId string
	ApiKey   string
	MerchId  string
}

func NewStripe() PaymentProcessor {
	return &Stripe{
		ClientId: "33334",
		MerchId:  "345",
		ApiKey:   "222",
	}
}

func (stripe *Stripe) Pay(amount float64) (string, error) {
	if amount < 0 {
		return "", utils.CustomeErrorFuncStripe("amount can not be negative")
	} else {
		trns := strconv.Itoa(rand.Int())
		obj := storage.NewTransaction(trns, "paypal")
		obj.AddTransaction(trns, "paypal")

		return "transaction Successull using stripe ", nil

	}

}

func (stripe *Stripe) Refund(transactionId string) (string, error) {

	if transactionId == "" {
		return "", utils.CustomeErrorFuncStripe("transaction id invalid")
	} else {
		for _, x := range storage.List {
			if x.TransactionId == transactionId {
				fmt.Println(x)
				return "payment processing refund using stripe", nil

			}
		}
	}
	return "", utils.CustomeErrorFuncStripe("transactionid not found")
}

func (stripe *Stripe) GetProviderName() (string, error) {
	//fmt.Println("paypal")
	//for _, x := range storage.List {
	//	if x.TransactionId == transactionID {
	//		fmt.Println("payment processing refund")
	//		fmt.Println(x)
	//	}
	//}
	return "stripe", nil
}
