package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	query := `INSERT INTO users (first_name, last_name, birth_date, gender, interests, city, username, password)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := DB.Exec(query, u.FirstName, u.LastName, u.BirthDate, u.Gender, u.Interests, u.City, u.Username, u.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	var dbPass string
	err := DB.QueryRow("SELECT password FROM users WHERE username=$1", input.Username).Scan(&dbPass)
	if err != nil || dbPass != input.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var u User
	err = DB.QueryRow("SELECT id, first_name, last_name, birth_date, gender, interests, city, username FROM users WHERE id=$1", id).
		Scan(&u.ID, &u.FirstName, &u.LastName, &u.BirthDate, &u.Gender, &u.Interests, &u.City, &u.Username)

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)
}
