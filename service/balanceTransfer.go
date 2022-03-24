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

func (d *DataBase) BalanceTransfer(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	if r.Method == http.MethodGet {

		url := r.URL.Query()

		rateId := url.Get("id")
		id, _ := strconv.Atoi(rateId)

		rateBalance := url.Get("balance")
		balance, _ := strconv.ParseFloat(rateBalance, 64)

		rateToId := url.Get("toid")
		id2, _ := strconv.Atoi(rateToId)

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

		row := tx.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)
		var balanceCheck float64

		err = row.Scan(&balanceCheck)
		fmt.Println("****************InCheck", balance, "****************************")
		if err != nil || balanceCheck < 0 {
			w.Write([]byte("insufficient funds"))
			w.WriteHeader(400)
			tx.Rollback()
			return
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", balance, id2)
		if err != nil {
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

		fmt.Println("*******************BalanceTRANSFER", user.Id, user.ToId, user.Balance, "****************")

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
		fmt.Println("****************InCheck", balance, "****************************")
		if err != nil || balance < 0 {
			w.Write([]byte("insufficient funds"))
			w.WriteHeader(400)
			tx.Rollback()
			return
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", user.Balance, user.ToId)
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(777, err)
			return
		}

		balanceOut := strconv.Itoa(int(user.Balance))
		w.Write([]byte("Transfer" + " " + balanceOut + "p"))
		return
	}

}
