package app

import (
	"avito/internal/domain"
	"avito/internal/handler"
	"avito/internal/repository"
	"database/sql"
)

type UserComposites struct {
	Storage domain.Storage
	Service repository.Service
	Handler repository.Handler
}

func Composite(db *sql.DB) (*UserComposites, error) {
	// Init Storage
	repository := repository.NewRepository(db)
	// Init Service
	service := domain.NewService(repository)
	// Init Handler
	handler := handler.NewHandler(service)

	return &UserComposites{
		Storage: repository,
		Service: service,
		Handler: handler,
	}, nil
}
