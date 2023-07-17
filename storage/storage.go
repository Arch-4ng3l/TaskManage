package storage

import (
	"github.com/Arch-4ng3l/TaskManage/types"
)

type Storage interface {
	AddNewAccount(*types.Account) error
	RemoveAccount(*types.Account) error
	GetAccountByName(string) (*types.Account, error)
	GetAccountByEmail(string) (*types.Account, error)
	AddNewTask(*types.Task) error
	RemoveTask(*types.Task) error
	TaskFromUser(string, string) (*types.Task, error)
	AllTasksFromUser(string) ([]*types.Task, error)
}
