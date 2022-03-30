package domain

import (
	"avito/internal/model"
)

func (s *Service) CreateUser(action string, user *model.User) ([]byte, error) {

	resp := make(map[string]string)

	err := s.Storage.AddUser(user)
	if err != nil {
		return nil, err
	}

	jsonResp, err := PrepareJson(action, resp, user)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}
