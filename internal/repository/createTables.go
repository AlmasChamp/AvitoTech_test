package repository

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,        
		email VARCHAR NOT NULL,       
		password VARCHAR NOT NULL,
		uuid VARCHAR NOT NULL,
		balance DECIMAL DEFAULT 0
	  );
	`)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Table Successfully Create!")
	return nil
}
