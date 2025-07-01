package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := os.Getenv("DB_CONN")
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("DB ping error: ", err)
	}
}
