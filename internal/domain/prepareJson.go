package domain

import (
	"avito/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func PrepareJson(flag string, resp map[string]string, user *model.User) ([]byte, error) {

	if flag == "reg" {
		resp["Id"] = strconv.Itoa(user.Id)
		resp["Email"] = user.Email
		resp["Password"] = user.Password
	} else if flag == "balance" {
		balance := fmt.Sprint(user.Balance)
		resp["Id"] = strconv.Itoa(user.Id)
		resp["Balance"] = balance
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return jsonResp, err
	}
	return jsonResp, err
}
