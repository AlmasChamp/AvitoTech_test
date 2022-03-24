package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func (d *DataBase) AddUser(w http.ResponseWriter, r *http.Request) {

	user := &User{
		Uuid: uuid.NewV1().String(),
	}

	fmt.Println("#################", user.Email, user.Password, "########################")

	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Body doesn't have any data")
		}

		if err = json.Unmarshal(body, user); err != nil {
			log.Println("JSON data isn't correct")
		}

		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~", user.Email, user.Password, user.Uuid, "~~~~~~~~~~~~~~~~~~~~~~~~")

		_, err = d.Db.Exec(`INSERT INTO users (email,password,uuid)
		VALUES ($1,$2,$3)`, user.Email, user.Password, user.Uuid)
		if err != nil {
			log.Println(err)
			return
		}

		row := d.Db.QueryRow("SELECT id FROM users WHERE email= $1", user.Email)

		err = row.Scan(&user.Id)

		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["userEmail"] = user.Email
		resp["userPassword"] = user.Password
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}
