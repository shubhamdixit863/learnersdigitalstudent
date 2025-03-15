package main

import (
	"fmt"
	"session8/practical/internal/services"
)

func main() {
	ps := services.NewPaymentServices()
  paypal := &services.PayPal{ClientID: "paypal123",APIKey: "pay_ApI",MerchantID: "abc"}
	strip:= &services.Stripe{ClientID: "stripe234",APIKey: "stripe_API",MerchantID: "ABC"}
	razorpay:=&services.Razorpay{ClientID: "Razorpay123",APIKey: "razor_API",MerchantID: "xyz"}


 ps.RegisterProvider(paypal)
 ps.RegisterProvider(strip)
 ps.RegisterProvider(razorpay)

 payment,err:=ps.ProcessPayment("PayPal",7800)
 if err!=nil{
	fmt.Println(err)
 }
 fmt.Println(payment)


 refund,err:=ps.IssueRefund("Stripe","TX-123")
 if err!=nil{
	fmt.Println(err)
 }
 fmt.Println(refund)

pro := ps.Providers["PayPal"]

services.Exatractdetails(pro)

}