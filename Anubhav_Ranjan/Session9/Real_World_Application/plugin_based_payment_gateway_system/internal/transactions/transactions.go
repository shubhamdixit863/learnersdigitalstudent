package transactions

import "fmt"

type Transaction struct {
	TransactionID string
	Amount        float64
	ProviderName  string
	Status        string
}

func (t Transaction) String() string {
	return fmt.Sprintf("Transaction ID: %s, Amount: %.2f, Provider: %s, Status: %s",
		t.TransactionID, t.Amount, t.ProviderName, t.Status)
}
