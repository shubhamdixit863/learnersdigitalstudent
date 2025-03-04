package services

import "payment-gateway/internal/util"

type Transaction struct {
	ID       string
	Provider string
	Amount   float64
	Status   string
}

type TransactionService struct {
	Transaction map[string]Transaction
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		Transaction: make(map[string]Transaction),
	}
}

func (ts *TransactionService) CreateTransaction(provider string, amount float64) string {
	transactionID := util.GetRandomUUID()
	ts.Transaction[transactionID] = Transaction{
		ID: 	 transactionID,
		Provider: provider,
		Amount:  amount,
		Status:  "PENDING",
	}
	return transactionID
}

func (ts *TransactionService) GetTransaction(transactionID string) (Transaction, bool) {
	transaction, ok := ts.Transaction[transactionID]
	return transaction, ok
}

func (ts *TransactionService) UpdateTransactionStatus(transactionID string, status string) bool {
	if transaction, ok := ts.Transaction[transactionID]; ok {
		transaction.Status = status
		ts.Transaction[transactionID] = transaction
		return true
	}
	return false
}

func (ts *TransactionService) GetAllTransactions() map[string]Transaction {
	return ts.Transaction
}