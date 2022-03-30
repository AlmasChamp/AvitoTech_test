package handler

import (
	"avito/internal/model"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) Balance(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}

	user := &model.User{}

	url := r.URL.Query()

	rateId := url.Get("id")

	if rateId == "" {
		w.WriteHeader(400)
		w.Write([]byte("id isn't correct"))
		return
	}
	user.Id, _ = strconv.Atoi(rateId)

	action := "balance"

	jsonResp, err := h.Service.ShowBalance(action, user)

	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResp)

	return

}
