package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type DataBase struct {
	db *sql.DB
}

type User struct {
	Id       int     `json:"id"`
	ToId     int     `json:"id2"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Uuid     string  `json:"uuid"`
	Balance  float64 `json:"balance"`
}

func main() {

	const (
		host     = "db"
		port     = "5432"
		user     = "postgres"
		password = "12345"
		dbname   = "postgres"
		sslMode  = "disable"
	)

	// connStr := "user=postgres password=12345 dbname=postgres sslmode=disable"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DataBase := &DataBase{
		db: db,
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", DataBase.AddUser)
	mux.HandleFunc("/balanceUser", DataBase.Balance)
	mux.HandleFunc("/balanceUp", DataBase.BalanceIncrease)
	mux.HandleFunc("/balanceDown", DataBase.BalanceDecrease)
	mux.HandleFunc("/balanceTransfer", DataBase.BalanceTransfer)
	// decrease
	fmt.Println("Start on port 8080")
	http.ListenAndServe(":8080", mux)
}

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

		_, err = d.db.Exec(`INSERT INTO users (email,password,uuid)
		VALUES ($1,$2,$3)`, user.Email, user.Password, user.Uuid)
		if err != nil {
			log.Println(err)
			return
		}

		row := d.db.QueryRow("SELECT id FROM users WHERE email= $1", user.Email)

		err = row.Scan(&user.Id)

		if err != nil {
			log.Println(err)
			return
		}

		userId := strconv.Itoa(user.Id)
		w.Write([]byte("Your id is" + " " + userId))
		return
	}
}

func (d *DataBase) BalanceIncrease(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	if r.Method == http.MethodGet {
		url := r.URL.Query()
		rateId := url.Get("id")
		id, _ := strconv.Atoi(rateId)

		rateBalance := url.Get("balance")
		balance, _ := strconv.ParseFloat(rateBalance, 64)
		fmt.Println("!!!!!!!!!!!!!!!!!!!PARSE STRING INTO FLOAT64!!!!!!!!!!!!!!!!!!!!!!", id, balance)
		_, err := d.db.Exec(`UPDATE users
		SET balance = $1+balance
		WHERE id = $2;`, balance, id)

		if err != nil {
			log.Println(err)
			return
		}
		balanceOut := strconv.Itoa(int(balance))
		w.Write([]byte("added" + " " + balanceOut + "p"))
		return
	}

	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("")
			return
		}

		if err = json.Unmarshal(body, user); err != nil {
			log.Println("JSON data isn't correct")
		}

		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~BalanceIncrease", user.Email, user.Password, user.Uuid, user.Balance, "~~~~~~~~~~~~~~~~~~~~~~~~")

		_, err = d.db.Exec(`UPDATE users
		SET balance = $1+balance
		WHERE id = $2;`, user.Balance, user.Id)

		if err != nil {
			log.Println(err)
			return
		}

	}
	balance := strconv.Itoa(int(user.Balance))
	w.Write([]byte("added" + " " + balance + "p"))
	return
}

func (d *DataBase) BalanceDecrease(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	if r.Method == http.MethodGet {

		url := r.URL.Query()

		rateId := url.Get("id")
		id, _ := strconv.Atoi(rateId)

		rateBalance := url.Get("balance")
		balance, _ := strconv.ParseFloat(rateBalance, 64)

		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, time.Second*5)

		tx, err := d.db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", balance, id)
		if err != nil {
			tx.Rollback()
			return
		}

		row := tx.QueryRow("SELECT balance FROM users WHERE id= $1", id)
		var balanceCheck float64

		err = row.Scan(&balanceCheck)

		if err != nil || balanceCheck < 0 {
			w.WriteHeader(400)
			w.Write([]byte("insufficient funds"))
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(777, err)
			return
		}
		balanceOut := strconv.Itoa(int(balance))
		w.Write([]byte(balanceOut + "p"))
		return
	}

	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)

		if err = json.Unmarshal(body, user); err != nil {
			log.Println("JSON data isn't correct")
		}

		fmt.Println("*******************BalanceDecrease", user.Id, user.Balance, "****************")

		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, time.Second*5)

		tx, err := d.db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", user.Balance, user.Id)
		if err != nil {
			tx.Rollback()
			return
		}

		row := tx.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)
		var balance float64

		err = row.Scan(&balance)

		if err != nil || balance < 0 {
			w.WriteHeader(400)
			w.Write([]byte("insufficient funds"))
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(777, err)
			return
		}

		balanceOut := strconv.Itoa(int(balance))
		w.Write([]byte(balanceOut + "p"))
		return
	}

}

func (d *DataBase) BalanceTransfer(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	if r.Method == http.MethodGet {

		url := r.URL.Query()

		rateId := url.Get("id")
		id, _ := strconv.Atoi(rateId)

		rateBalance := url.Get("balance")
		balance, _ := strconv.ParseFloat(rateBalance, 64)

		rateToId := url.Get("toid")
		id2, _ := strconv.Atoi(rateToId)

		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, time.Second*5)

		tx, err := d.db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", balance, id)
		if err != nil {
			tx.Rollback()
			return
		}

		row := tx.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)
		var balanceCheck float64

		err = row.Scan(&balanceCheck)
		fmt.Println("****************InCheck", balance, "****************************")
		if err != nil || balanceCheck < 0 {
			w.Write([]byte("insufficient funds"))
			w.WriteHeader(400)
			tx.Rollback()
			return
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", balance, id2)
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(777, err)
			return
		}
		balanceOut := strconv.Itoa(int(balance))
		w.Write([]byte(balanceOut + "p"))
		return
	}

	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)

		if err = json.Unmarshal(body, user); err != nil {
			log.Println("JSON data isn't correct")
		}

		fmt.Println("*******************BalanceTRANSFER", user.Id, user.ToId, user.Balance, "****************")

		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, time.Second*5)

		tx, err := d.db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", user.Balance, user.Id)
		if err != nil {
			tx.Rollback()
			return
		}

		row := tx.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)
		var balance float64

		err = row.Scan(&balance)
		fmt.Println("****************InCheck", balance, "****************************")
		if err != nil || balance < 0 {
			w.Write([]byte("insufficient funds"))
			w.WriteHeader(400)
			tx.Rollback()
			return
		}

		_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", user.Balance, user.ToId)
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(777, err)
			return
		}

		balanceOut := strconv.Itoa(int(user.Balance))
		w.Write([]byte("Transfer" + " " + balanceOut + "p"))
		return
	}

}

func (d *DataBase) Balance(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	if r.Method == http.MethodGet {

		url := r.URL.Query()

		rateId := url.Get("id")
		id, _ := strconv.Atoi(rateId)

		row := d.db.QueryRow("SELECT balance FROM users WHERE id= $1", id)

		err := row.Scan(&user.Balance)

		if err != nil {
			log.Println(err)
			return
		}
		balanceOut := strconv.Itoa(int(user.Balance))
		w.Write([]byte("Your" + " " + "balance" + " " + balanceOut + "p"))
		return
	}

	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("")
			return
		}

		if err = json.Unmarshal(body, user); err != nil {
			log.Println("JSON data isn't correct")
		}

		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~Balance", user.Id, "~~~~~~~~~~~~~~~~~~~~~~~~")

		row := d.db.QueryRow("SELECT balance FROM users WHERE id= $1", user.Id)

		err = row.Scan(&user.Balance)

		if err != nil {
			log.Println(err)
			return
		}
		balance := strconv.Itoa(int(user.Balance))
		w.Write([]byte("Your" + " " + "balance" + " " + balance + "p"))
		return
	}
}
