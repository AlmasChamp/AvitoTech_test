package service

import "database/sql"

type DataBase struct {
	Db *sql.DB
}

type User struct {
	Id       int     `json:"id"`
	ToId     int     `json:"id2"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Uuid     string  `json:"uuid"`
	Balance  float64 `json:"balance"`
}
