package handler

import (
	"avito/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) BalanceDecrease(w http.ResponseWriter, r *http.Request) {

	user := &model.User{}

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err = json.Unmarshal(body, user); err != nil {
		log.Println("JSON data isn't correct")
	}

	if user.Id == 0 || user.Balance == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Id or Balance is incorrect "))
		return
	}

	action := "down"

	jsonResp, err := h.Service.BalanceDecrs(action, user)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResp)
}
