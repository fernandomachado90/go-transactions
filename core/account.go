package core

type Account struct {
	ID             int
	DocumentNumber string
}

type AccountManager struct {
	db interface {
		CreateAccount(Account) (Account, error)
	}
}

func (m *AccountManager) Create(account Account) (Account, error) {
	account, err := m.db.CreateAccount(account)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}
