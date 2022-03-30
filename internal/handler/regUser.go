package handler

import (
	"avito/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

//
func (h *Handler) RegUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	user := &model.User{
		Uuid: uuid.NewV1().String(),
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		log.Println("Body doesn't have any data")
	}

	if err = json.Unmarshal(body, user); err != nil {
		log.Println("JSON data isn't correct")
	}

	if user.Email == "" || user.Password == "" {
		w.WriteHeader(400)
		w.Write([]byte("Email or Password is incorrect "))
		return
	}

	action := "reg"
	jsonResp, err := h.Service.CreateUser(action, user)

	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResp)
	return

}
