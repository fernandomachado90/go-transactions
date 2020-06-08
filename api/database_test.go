package main

import (
	"github.com/fernandomachado90/go-transactions/core"
	"github.com/stretchr/testify/mock"
)

type dbMock struct {
	mock.Mock
}

func (db *dbMock) CreateAccount(account core.Account) (core.Account, error) {
	args := db.Called(account)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return core.Account{}, err
	}

	account.ID = id
	return account, nil
}

func (db *dbMock) FindAccount(id int) (core.Account, error) {
	args := db.Called(id)
	doc := args.String(0)
	err := args.Error(1)
	if err != nil {
		return core.Account{}, err
	}

	return core.Account{ID: id, DocumentNumber: doc}, nil
}

func (db *dbMock) CreateTransaction(transaction core.Transaction) (core.Transaction, error) {
	args := db.Called(transaction)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return core.Transaction{}, err
	}

	transaction.ID = id
	return transaction, nil
}

func (db *dbMock) FindOperation(id int) (core.Operation, error) {
	args := db.Called(id)
	credit := args.Bool(0)
	err := args.Error(1)
	if err != nil {
		return core.Operation{}, err
	}

	return core.Operation{ID: id, Credit: credit}, nil
}
