package handler

import (
	"avito/internal/model"
	"encoding/json"
	"fmt"
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

	fmt.Println("#######Regpage##########", user.Email, user.Password, "###########Before###########")

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

	fmt.Println("~~~~~~~~~Regpage~~~~~~~~~~~~~~~", user.Email, user.Password, user.Uuid, "~~~~~~~~~~~~~After~~~~~~~~~~~")
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	flag := "reg"
	jsonResp, err := h.Service.CreateUser(flag, user)

	// if err = service.CreateUser(d.Db, user); err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(500)
	// 	return
	// }

	if err != nil {
		log.Println(err)
		return
	}

	// resp := make(map[string]string)
	// flag := "reg"

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// jsonResp, err := PrepareJson(flag, resp, user)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(500)
	// 	w.Write([]byte("Server Error "))
	// 	return
	// }

	w.Write(jsonResp)
	return

}
