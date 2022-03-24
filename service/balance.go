package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (d *DataBase) Balance(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	if r.Method == http.MethodGet {

		url := r.URL.Query()

		rateId := url.Get("id")
		id, _ := strconv.Atoi(rateId)

		row := d.Db.QueryRow("SELECT balance FROM users WHERE id= $1", id)

		err := row.Scan(&user.Balance)

		if err != nil {
			log.Println(err)
			return
		}
		balanceOut := strconv.Itoa(int(user.Balance))
		w.Write([]byte("Your" + " " + "balance" + " " + balanceOut + "p"))
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

		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~Balance", user.Id, "~~~~~~~~~~~~~~~~~~~~~~~~")

		row := d.Db.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)

		err = row.Scan(&user.Balance)

		if err != nil {
			log.Println(err)
			return
		}
		balance := strconv.Itoa(int(user.Balance))
		w.Write([]byte("Your" + " " + "balance" + " " + balance + "p"))
		return
	}
}
