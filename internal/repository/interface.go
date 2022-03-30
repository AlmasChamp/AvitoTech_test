package repository

import (
	"avito/internal/model"
	"net/http"
)

type Service interface {
	CreateUser(action string, user *model.User) ([]byte, error)
	ShowBalance(action string, user *model.User) ([]byte, error)
	BalanceIncrs(action string, user *model.User) ([]byte, error)
	BalanceDecrs(action string, user *model.User) ([]byte, error)
	Transfer(action string, user *model.User) ([]byte, error)
}

type Handler interface {
	Register(mux *http.ServeMux)
}
