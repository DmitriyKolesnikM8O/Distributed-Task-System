package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	model "github.com/KolesnikM8O/distributed-task-system/auth-service/internal/service/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupReq model.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&signupReq)
	if err != nil {
		log.Printf("Error decoding signup request: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if signupReq.Username == "" || signupReq.Email == "" || signupReq.Password == "" {
		log.Printf("Invalid signup request")
		http.Error(w, "Invalid signup request", http.StatusBadRequest)
		return
	}

	var existingUsername string
	err = s.db.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1", signupReq.Username).Scan(&existingUsername)
	if err == nil {
		log.Printf("Username already taken")
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupReq.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	role := "user"
	_, err = s.db.Exec(context.Background(),
		"INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)",
		signupReq.Username, signupReq.Email, hashedPassword, role)
	if err != nil {
		log.Printf("Error inserting user into DB: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User %s signed up", signupReq.Username)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
