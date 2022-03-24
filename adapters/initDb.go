package adapters

import (
	"avito/service"
	"database/sql"
	"fmt"
	"log"
)

func InitDb() *service.DataBase {

	const (
		host     = "db"
		port     = "5432"
		user     = "postgres"
		password = "12345"
		dbname   = "postgres"
		sslMode  = "disable"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")

	_, err = db.Exec(`
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
	}

	log.Println("Table Successfully Create!")

	return &service.DataBase{
		Db: db,
	}
}
