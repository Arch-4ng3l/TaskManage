package main

type Storage interface {
	AddNewAccount(*Account) error
	RemoveAccount(*Account) error
	GetAccountByName(string) (*Account, error)
	GetAccountByEmail(string) (*Account, error)
}
