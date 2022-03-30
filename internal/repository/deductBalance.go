package repository

import (
	"avito/internal/model"
	"context"
	"log"
	"time"
)

func (r *Repository) DeductBalance(user *model.User) error {

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*5)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	row := tx.QueryRow("SELECT uuid FROM users WHERE id= $1", user.Id)

	err = row.Scan(&user.Uuid)

	if err != nil || len(user.Uuid) <= 0 {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE uuid = $2", user.Balance, user.Uuid)
	if err != nil {
		tx.Rollback()
		return err
	}

	row = tx.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)
	var balance float64

	err = row.Scan(&balance)

	if err != nil || balance < 0 {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(777, err)
		return err
	}
	return nil
}
