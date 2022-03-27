package handler

import (
	"avito/internal/model"
	"fmt"
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

	fmt.Println(r.Method, "~~~~~~~~~~~~~~~~~~~~~~~~Balance", user.Id, "~~~~~~~~~~~~~~~~~~~~~~~~")

	url := r.URL.Query()

	rateId := url.Get("id")
	if rateId == "" {
		w.WriteHeader(400)
		w.Write([]byte("id isn't correct"))
		return
	}
	id, _ := strconv.Atoi(rateId)
	user.Id = id

	// row := d.Db.QueryRow("SELECT balance FROM users WHERE id= $1", id)

	// err := row.Scan(&user.Balance)

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// resp := make(map[string]string)

	action := "balance"

	jsonResp, err := h.Service.ShowBalance(action, user, id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
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

	return

}
