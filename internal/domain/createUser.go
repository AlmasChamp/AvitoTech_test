package domain

import (
	"avito/internal/model"
)

func (s *Service) CreateUser(flag string, user *model.User) ([]byte, error) {
	// if err := CanRegister(user); err != nil {
	// 	return err
	// }
	resp := make(map[string]string)

	err := s.Storage.AddUser(user)
	return nil, err

	jsonResp, err := PrepareJson(flag, resp, user)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}
