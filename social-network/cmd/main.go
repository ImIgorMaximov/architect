package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/user/register", RegisterUser).Methods("POST")
	r.HandleFunc("/user/get/{id}", GetUser).Methods("GET")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
