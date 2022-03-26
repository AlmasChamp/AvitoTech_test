package main

import (
	"avito/internal/app"
	"avito/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	db := repository.InitDb()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("close database: %v\n", err)
		}
	}(db)

	if err := repository.CreateTables(db); err != nil {
		log.Fatal(err)
		return
	}

	mux := http.NewServeMux()

	userComposite, err := app.Composites(db)
	if err != nil {
		log.Println(err)
		return
	}
	userComposite.Handler.Register(mux)

	fmt.Println("Start on port 9000")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
