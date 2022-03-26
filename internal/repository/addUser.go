package repository

import "avito/internal/model"

func (r *Repository) AddUser(user *model.User) error {

	_, err := r.db.Exec(`INSERT INTO users (email,password,uuid)
		VALUES ($1,$2,$3)`, user.Email, user.Password, user.Uuid)
	if err != nil {
		return err
	}

	row := r.db.QueryRow("SELECT id FROM users WHERE email= $1", user.Email)

	if err = row.Scan(&user.Id); err != nil {
		return err
	}

	return nil
}
