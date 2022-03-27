package handler

import (
	"avito/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) BalanceIncrease(w http.ResponseWriter, r *http.Request) {

	user := &model.User{}

	// if r.Method == http.MethodGet {
	// 	url := r.URL.Query()
	// 	rateId := url.Get("id")
	// 	id, _ := strconv.Atoi(rateId)

	// 	rateBalance := url.Get("balance")
	// 	balance, _ := strconv.ParseFloat(rateBalance, 64)
	// 	fmt.Println("!!!!!!!!!!!!!!!!!!!PARSE STRING INTO FLOAT64!!!!!!!!!!!!!!!!!!!!!!", id, balance)
	// 	_, err := d.Db.Exec(`UPDATE users
	// 	SET balance = $1+balance
	// 	WHERE id = $2;`, balance, id)

	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	balanceOut := strconv.Itoa(int(balance))
	// 	w.Write([]byte("added" + " " + balanceOut + "p"))
	// 	return
	// }

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("JSON data isn't correct")
		return
	}

	if err = json.Unmarshal(body, user); err != nil {
		log.Println("JSON data isn't correct")
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~BalanceIncrease", user.Id, user.Balance, "~~~~~~~~~~~~~~~~~~~~~~~~")

	// _, err = d.Db.Exec(`UPDATE users
	// SET balance = $1+balance
	// WHERE id = $2;`, user.Balance, user.Id)

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	action := "Up"

	jsonResp, err := h.Service.BalanceIncrs(action, user)

	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// jsonResp, err := prepareJson(flag, resp, user)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(500)
	// 	w.Write([]byte("Server Error "))
	// 	return
	// }

	w.Write(jsonResp)

	// balance := strconv.Itoa(int(user.Balance))
	// w.Write([]byte("added" + " " + balance + "p"))
	// return
}
