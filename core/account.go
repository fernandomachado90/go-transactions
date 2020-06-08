package core

import "errors"

type Account struct {
	ID             int
	DocumentNumber string
}

type AccountManager struct {
	db
}

type db interface {
	CreateAccount(Account) (Account, error)
	FindAccount(id int) (Account, error)
}

func NewAccountManager(db db) *AccountManager {
	return &AccountManager{db}
}

func (m *AccountManager) Create(account Account) (Account, error) {
	if account.DocumentNumber == "" {
		return Account{}, errors.New("missing required attribute")
	}

	account, err := m.db.CreateAccount(account)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

func (m *AccountManager) Find(id int) (Account, error) {
	account, err := m.db.FindAccount(id)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}
