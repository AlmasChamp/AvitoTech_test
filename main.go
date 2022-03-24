package main

import (
	"avito/adapters"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	DataBase := adapters.InitDb()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("close database: %v\n", err)
		}
	}(DataBase.Db)

	mux := http.NewServeMux()
	mux.HandleFunc("/", DataBase.AddUser)
	mux.HandleFunc("/balanceUser", DataBase.Balance)
	mux.HandleFunc("/balanceUp", DataBase.BalanceIncrease)
	mux.HandleFunc("/balanceDown", DataBase.BalanceDecrease)
	mux.HandleFunc("/balanceTransfer", DataBase.BalanceTransfer)
	fmt.Println("Start on port 9000")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
