package main

import (
	"github.com/fernandomachado90/go-transactions/core"
	"github.com/stretchr/testify/mock"
)

type dbMock struct {
	mock.Mock
}

func (m *dbMock) CreateAccount(account core.Account) (core.Account, error) {
	args := m.Called(account)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return core.Account{}, err
	}

	account.ID = id
	return account, nil
}

func (m *dbMock) FindAccount(id int) (core.Account, error) {
	args := m.Called(id)
	doc := args.String(0)
	err := args.Error(1)
	if err != nil {
		return core.Account{}, err
	}

	return core.Account{ID: id, DocumentNumber: doc}, nil
}

func (m *dbMock) CreateTransaction(transaction core.Transaction) (core.Transaction, error) {
	args := m.Called(transaction)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return core.Transaction{}, err
	}

	transaction.ID = id
	return transaction, nil
}
