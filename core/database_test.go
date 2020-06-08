package core

import "github.com/stretchr/testify/mock"

type dbMock struct {
	mock.Mock
}

func (m *dbMock) CreateAccount(account Account) (Account, error) {
	args := m.Called(account)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return Account{}, err
	}

	account.ID = id
	return account, nil
}

func (m *dbMock) FindAccount(id int) (Account, error) {
	args := m.Called(id)
	doc := args.String(0)
	err := args.Error(1)
	if err != nil {
		return Account{}, err
	}

	return Account{ID: id, DocumentNumber: doc}, nil
}

func (m *dbMock) CreateTransaction(transaction Transaction) (Transaction, error) {
	args := m.Called(transaction)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return Transaction{}, err
	}

	transaction.ID = id
	return transaction, nil
}

func (m *dbMock) FindOperation(id int) (Operation, error) {
	args := m.Called(id)
	credit := args.Bool(0)
	err := args.Error(1)
	if err != nil {
		return Operation{}, err
	}

	return Operation{ID: id, Credit: credit}, nil
}
