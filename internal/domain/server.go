package domain

import "avito/internal/repository"

type Service struct {
	Storage Storage
}

func NewService(repository Storage) repository.Service {
	return &Service{Storage: repository}
}
