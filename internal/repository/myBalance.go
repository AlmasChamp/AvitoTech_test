package repository

import (
	"avito/internal/model"
	"fmt"
)

func (r *Repository) MyBalance(user *model.User, id int) error {

	row := r.db.QueryRow("SELECT id FROM users WHERE id= $1", id)

	err := row.Scan(&user.Id)
	fmt.Println(user.Id, "^^^^^^^^^^^^^^")
	if err != nil {
		return err
	}

	row = r.db.QueryRow("SELECT balance FROM users WHERE id= $1", id)

	err = row.Scan(&user.Balance)
	fmt.Println(user.Balance, "^^^^^^^^^^^^^^")
	if err != nil {
		return err
	}
	return nil
}
