package domain

import (
	"avito/internal/model"
	"log"
)

func (s *Service) ShowBalance(action string, user *model.User, id int) ([]byte, error) {

	resp := make(map[string]string)

	if err := s.Storage.MyBalance(user, id); err != nil {
		return nil, err
	}

	jsonResp, err := PrepareJson(action, resp, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return jsonResp, nil

}
