package core

import "time"

type Transaction struct {
	ID          int
	AccountID   int
	OperationID int
	Amount      float64
	EventDate   time.Time
}

type TransactionManager struct {
	db interface {
		CreateTransaction(Transaction) (Transaction, error)
	}
}

func (m *TransactionManager) Create(transaction Transaction) (Transaction, error) {
	transaction, err := m.db.CreateTransaction(transaction)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}
