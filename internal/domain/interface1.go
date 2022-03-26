package domain

import "avito/internal/model"

type Storage interface {
	AddUser(user *model.User) error
	// GetUserPassword(email string) (string, error)
	// GetUserId(email string) (string, error)
	// SetCookie(cookieVal string, cookieExp int, id string) error
	// DeleteCookie(cookie string) error
	// GetAllPosts() ([]entities.Post, error)
	// GetValueCookie(userCookie string) (string, error)
}
