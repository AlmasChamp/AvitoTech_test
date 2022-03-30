package model

type User struct {
	Id        int     `json:"id"`
	ReciverId int     `json:"reciver"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Uuid      string  `json:"uuid"`
	Balance   float64 `json:"balance"`
}
