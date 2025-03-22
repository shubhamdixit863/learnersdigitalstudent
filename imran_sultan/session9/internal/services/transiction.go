package services

import (
	"fmt"
	"time"
)

type Transiction struct {
	ClientId string
	Time     time.Time
	Amount   float64
	Transid  string
	Provider string
}

var Mydata []Transiction

func NewTransiction() *Transiction {
	return &Transiction{ClientId: "", Time: time.Now(), Amount: 3.6, Transid: ""}
}
func (d Transiction) AddNewTransiction(clientid string, amt float64, transid string, provid string) []Transiction {
	trs := Transiction{
		ClientId: clientid,
		Time:     time.Now(),
		Amount:   amt,
		Transid:  transid,
		Provider: provid,
	}
	Mydata = append(Mydata, trs)
	fmt.Println(Mydata)
	return Mydata
}
