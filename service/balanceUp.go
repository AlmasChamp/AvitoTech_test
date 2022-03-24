package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (d *DataBase) BalanceIncrease(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	if r.Method == http.MethodGet {
		url := r.URL.Query()
		rateId := url.Get("id")
		id, _ := strconv.Atoi(rateId)

		rateBalance := url.Get("balance")
		balance, _ := strconv.ParseFloat(rateBalance, 64)
		fmt.Println("!!!!!!!!!!!!!!!!!!!PARSE STRING INTO FLOAT64!!!!!!!!!!!!!!!!!!!!!!", id, balance)
		_, err := d.Db.Exec(`UPDATE users
		SET balance = $1+balance
		WHERE id = $2;`, balance, id)

		if err != nil {
			log.Println(err)
			return
		}
		balanceOut := strconv.Itoa(int(balance))
		w.Write([]byte("added" + " " + balanceOut + "p"))
		return
	}

	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("")
			return
		}

		if err = json.Unmarshal(body, user); err != nil {
			log.Println("JSON data isn't correct")
		}

		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~BalanceIncrease", user.Email, user.Password, user.Uuid, user.Balance, "~~~~~~~~~~~~~~~~~~~~~~~~")

		_, err = d.Db.Exec(`UPDATE users
		SET balance = $1+balance
		WHERE id = $2;`, user.Balance, user.Id)

		if err != nil {
			log.Println(err)
			return
		}

	}
	balance := strconv.Itoa(int(user.Balance))
	w.Write([]byte("added" + " " + balance + "p"))
	return
}
