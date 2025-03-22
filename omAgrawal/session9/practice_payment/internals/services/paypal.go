package services

import (
	"fmt"
	"math/rand"
	"practice_payment/internals/storage"
	"practice_payment/internals/utils"
	"strconv"
	//"nexturn_Go/services"
)

type Paypal struct {
	ClientId      string
	ApiKey        string
	MerchId       string
	TranscationId string
}

func NewPaypal(client string, api string, merc string) PaymentProcessor {

	//fmt.Println(" Paypal initiated ")
	return &Paypal{
		ClientId: client,
		ApiKey:   api,
		MerchId:  merc,
	}

}

func (paypal *Paypal) Pay(amount float64) (string, error) {

	if amount < 0 {
		return "", utils.CustomeErrorFuncPaypal("amount is in negative")
	} else {
		trns := strconv.Itoa(rand.Int())
		obj := storage.NewTransaction(trns, "paypal")
		obj.AddTransaction(trns, "paypal")
		return "Paid using paypal ", nil

	}

}

func (paypal *Paypal) Refund(transactionId string) (string, error) {

	//fmt.Println("refunded", transactionId)
	if transactionId == "" {
		return "", utils.CustomeErrorFuncPaypal("transaction id invalid")
	} else {
		for _, x := range storage.List {
			if x.TransactionId == transactionId {
				fmt.Println(x)
				return "payment processing refund", nil

			}
		}
	}
	return "", utils.CustomeErrorFuncPaypal("transactionid not found")
}

func (paypal *Paypal) GetProviderName() (string, error) {
	return "Paypal", nil
}
