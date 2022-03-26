package repository

import (
	"avito/internal/model"
	"net/http"
)

type Service interface {
	CreateUser(flag string, user *model.User) ([]byte, error)
	// LogInUser(email, password string) (*http.Cookie, error)
	// LogOut(cookieVal string) *http.Cookie
	// AllPost() ([]entities.Post, error)
	// ValueCookie(userCookie string) (string, error)
}

type Handler interface {
	Register(mux *http.ServeMux)
}
