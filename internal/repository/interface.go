package repository

import (
	"avito/internal/model"
	"net/http"
)

type Service interface {
	CreateUser(action string, user *model.User) ([]byte, error)
	ShowBalance(action string, user *model.User, id int) ([]byte, error)
	BalanceIncrs(action string, user *model.User) ([]byte, error)
	BalanceDecrs(action string, user *model.User, id int) ([]byte, error)
	Transfer(action string, user *model.User) ([]byte, error)
	// LogInUser(email, password string) (*http.Cookie, error)
	// LogOut(cookieVal string) *http.Cookie
	// AllPost() ([]entities.Post, error)
	// ValueCookie(userCookie string) (string, error)
}

type Handler interface {
	Register(mux *http.ServeMux)
}
