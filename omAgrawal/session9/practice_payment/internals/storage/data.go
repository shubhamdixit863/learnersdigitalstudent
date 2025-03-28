package storage

import (
	"fmt"
	"time"
)

type Transaction struct {
	TransactionId string
	Time          time.Time
	Provider      string
}

var List = make([]Transaction, 0)

func NewTransaction(trns string, method string) *Transaction {
	return &Transaction{
		TransactionId: trns,
		Time:          time.Now(),
		Provider:      method,
	}
}

func (d Transaction) AddTransaction(transId string, provide string) {

	List = append(List, d)
	fmt.Println("Added transaction ", List)

}
