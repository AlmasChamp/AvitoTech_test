package handler

import (
	"avito/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) BalanceTransfer(w http.ResponseWriter, r *http.Request) {

	user := &model.User{}

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err = json.Unmarshal(body, user); err != nil {
		log.Println("JSON data isn't correct")
	}

	if user.Id == 0 || user.Balance == 0 || user.ReciverId == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Id,Balance or ReciverId is incorrect "))
		return
	}

	fmt.Println(user.Id, user.Balance, user.ReciverId)

	action := "transfer"

	jsonResp, err := h.Service.Transfer(action, user)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResp)

	balanceOut := strconv.Itoa(int(user.Balance))
	w.Write([]byte("Transfer" + " " + balanceOut + "p"))
	return

}
