package services

import (
	"fmt"
	"session9/internal/util"
	// "session9/internal/services"
	// "session9/internal/services"
)

type PayPal struct {
	Email string
}

func NewPayPal() PaymentServices {

	fmt.Println("Welcome to PayPal Gateway!...")
	fmt.Println("Enter your Email: ")
	var email string
	fmt.Scanln(&email)
	return &PayPal{
		Email: email,
	}
}

func (p PayPal) Pay(amount float64) (string, error) {

	// randomNumber := rand.Intn(1000)
	// transid := strconv.Itoa(randomNumber)
	reqamount := 1000.0
	transid := "xyz"
	obj := NewTransiction()
	if amount < reqamount {
		return "", &util.PaypalError{Message: "Transiction failed not enough money!"}
	}

	obj.AddNewTransiction(p.Email, amount, transid, "PayPal")

	return fmt.Sprintln("You paid amount", amount, " paypal id: ", p.Email), nil
}
func (p PayPal) Refund(transactionID string) (string, error) {
	for i := 0; i < len(Mydata); i++ {
		if Mydata[i].Transid == transactionID {
			Mydata[i].Amount = 0
			return "Refunded successfully!", nil
		}
	}
	// fmt.Println()
	return "", &util.PaypalError{Message: "transiction not found!"}
}
func (p PayPal) GetProviderName() string {
	return "Payment done via PayPal!"
}
