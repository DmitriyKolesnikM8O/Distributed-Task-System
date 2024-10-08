package handlers

import (
	"context"
	"encoding/json"
	"log"

	"net/http"

	model "github.com/KolesnikM8O/distributed-task-system/auth-service/service/model"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	var loginReq model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var hashedPassword string
	err = db.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", loginReq.Username).Scan(&hashedPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginReq.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	log.Printf("User %s logged in", loginReq.Username)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SignupHandler(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	var signupReq model.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&signupReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if signupReq.Username == "" || signupReq.Email == "" || signupReq.Password == "" {

		http.Error(w, "Invalid signup request", http.StatusBadRequest)
		return
	}

	var existingUsername string
	err = db.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1", signupReq.Username).Scan(&existingUsername)
	if err == nil {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupReq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", signupReq.Username, signupReq.Email, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User %s signed up", signupReq.Username)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
