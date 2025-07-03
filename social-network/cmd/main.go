package main

import (
	"log"
	"net/http"

	"social-network/internal/db"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/login", db.Login).Methods("POST")
	r.HandleFunc("/user/register", db.RegisterUser).Methods("POST")
	r.HandleFunc("/user/get/{id}", db.GetUser).Methods("GET")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
