package transactions

import "fmt"

type Transaction struct {
	TranscationID   string
	ClientID        string
	MerchantID      string
	PaymentProvider string
}

var Transactions []Transaction

func ShowTransaction() {
	for i, t := range Transactions {
		fmt.Println("Transaction No.", i+1)
		fmt.Println("Provider Name:", t.PaymentProvider)
		fmt.Println("Merchant ID:", t.MerchantID)
		fmt.Println("Client ID:", t.ClientID)
		fmt.Println("Transaction ID:", t.TranscationID)
		fmt.Println()
	}
}
