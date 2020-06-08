package core

type db interface {
	CreateAccount(Account) (Account, error)
	FindAccount(id int) (Account, error)
	CreateTransaction(Transaction) (Transaction, error)
}
