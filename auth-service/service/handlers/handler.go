package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"net/http"

	model "github.com/KolesnikM8O/distributed-task-system/auth-service/service/model"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	var loginReq model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		log.Printf("Error decoding login request: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var hashedPassword string
	err = db.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", loginReq.Username).Scan(&hashedPassword)
	if err != nil {
		log.Printf("Error getting password from DB: %s", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginReq.Password))
	if err != nil {
		log.Printf("Error comparing passwords: %s", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	log.Printf("User %s logged in", loginReq.Username)
	var user model.User
	err = db.QueryRow(context.Background(), "SELECT id, role FROM users WHERE username = $1", loginReq.Username).Scan(&user.ID, &user.Role)
	if err != nil {
		log.Printf("Error getting user from DB: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
	})
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		log.Printf("Error signing token: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SignupHandler(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
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
	err = db.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1", signupReq.Username).Scan(&existingUsername)
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
	_, err = db.Exec(context.Background(),
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
