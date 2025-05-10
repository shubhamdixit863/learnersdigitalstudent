package services


// type Service interface{
// 	GetData() string
// 	SearchData() string
// 	AddData(data string) error
// }


type PaymentProcessor interface{
	Pay(amount float64) string
	Refund(transactionID string) string
	GetProviderName() string
}

