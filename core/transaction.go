package core

import (
	"errors"
	"time"
)

type Transaction struct {
	ID          int
	AccountID   int
	OperationID int
	Amount      float64
	EventDate   time.Time
}

type TransactionManager struct {
	db
}

func NewTransactionManager(db db) *TransactionManager {
	return &TransactionManager{db}
}

func (m *TransactionManager) Create(transaction Transaction) (Transaction, error) {
	if transaction.AccountID == 0 || transaction.OperationID == 0 {
		return Transaction{}, errors.New("missing required attribute")
	}
	transaction.EventDate = time.Now()

	transaction, err := m.db.CreateTransaction(transaction)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}
