package repository

import (
	"avito/internal/model"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *Repository) BalanceTransfer(user *model.User) error {

	var checkBalance float64
	var checkUuid string

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*5)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	row := tx.QueryRow("SELECT uuid FROM users WHERE id= $1", user.Id)

	err = row.Scan(&user.Uuid)

	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE uuid = $2", user.Balance, user.Uuid)
	if err != nil {
		tx.Rollback()
		return err
	}

	row = tx.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)

	err = row.Scan(&checkBalance)
	fmt.Println("****************InCheck", user.Balance, "****************************")
	if err != nil || user.Balance < 0 {
		tx.Rollback()
		return err
	}

	row = tx.QueryRow("SELECT uuid FROM users WHERE id= $1", user.ReciverId)

	err = row.Scan(&checkUuid)

	if err != nil || len(checkUuid) <= 0 {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance + $1 WHERE uuid = $2", user.Balance, checkUuid)
	if err != nil {
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
