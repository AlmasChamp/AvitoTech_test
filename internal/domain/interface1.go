package domain

import "avito/internal/model"

type Storage interface {
	AddUser(user *model.User) error
	MyBalance(user *model.User) error
	AddBalance(user *model.User) error
	DeductBalance(user *model.User) error
	BalanceTransfer(user *model.User) error
}
