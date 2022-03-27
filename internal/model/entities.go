package model

// type DataBase struct {
// 	Db *sql.DB
// }

// type Repository struct {
// 	db *sql.DB
// }

type User struct {
	Id        int     `json:"id"`
	ReciverId int     `json:"reciver"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Uuid      string  `json:"uuid"`
	Balance   float64 `json:"balance"`
}

// type User struct {
// 	Id        int     `json:"id"`
// 	ReciverId int     `json:"id2"`
// 	Email     string  `json:"email"`
// 	Password  string  `json:"password"`
// 	Uuid      string  `json:"uuid"`
// 	Balance   float64 `json:"balance"`
// }
