package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (d *DataBase) BalanceDecrease(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	if r.Method == http.MethodGet {

		url := r.URL.Query()

		rateId := url.Get("id")
		id, _ := strconv.Atoi(rateId)

		rateBalance := url.Get("balance")
		balance, _ := strconv.ParseFloat(rateBalance, 64)

		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, time.Second*5)

		tx, err := d.Db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", balance, id)
		if err != nil {
			tx.Rollback()
			return
		}

		row := tx.QueryRow("SELECT balance FROM users WHERE id= $1", id)
		var balanceCheck float64

		err = row.Scan(&balanceCheck)

		if err != nil || balanceCheck < 0 {
			w.WriteHeader(400)
			w.Write([]byte("insufficient funds"))
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(777, err)
			return
		}
		balanceOut := strconv.Itoa(int(balance))
		w.Write([]byte(balanceOut + "p"))
		return
	}

	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)

		if err = json.Unmarshal(body, user); err != nil {
			log.Println("JSON data isn't correct")
		}

		fmt.Println("*******************BalanceDecrease", user.Id, user.Balance, "****************")

		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, time.Second*5)

		tx, err := d.Db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", user.Balance, user.Id)
		if err != nil {
			tx.Rollback()
			return
		}

		row := tx.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)
		var balance float64

		err = row.Scan(&balance)

		if err != nil || balance < 0 {
			w.WriteHeader(400)
			w.Write([]byte("insufficient funds"))
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(777, err)
			return
		}

		balanceOut := strconv.Itoa(int(balance))
		w.Write([]byte(balanceOut + "p"))
		return
	}

}
