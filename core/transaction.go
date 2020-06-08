package core

import (
	"errors"
	"math"
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

func (m *TransactionManager) Create(t Transaction) (Transaction, error) {
	if t.AccountID == 0 || t.OperationID == 0 {
		return Transaction{}, errors.New("missing required attribute")
	}

	op, err := m.db.FindOperation(t.OperationID)
	if err != nil {
		return Transaction{}, errors.New("operation not found")
	}
	math.Abs(t.Amount)
	if !op.Credit {
		t.Amount = t.Amount * -1
	}

	t.EventDate = time.Now()

	t, err = m.db.CreateTransaction(t)
	if err != nil {
		return Transaction{}, err
	}

	return t, nil
}
