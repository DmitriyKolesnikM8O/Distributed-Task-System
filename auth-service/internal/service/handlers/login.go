package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KolesnikM8O/distributed-task-system/auth-service/internal/config/config"
	model "github.com/KolesnikM8O/distributed-task-system/auth-service/internal/service/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	cfg       = config.GetConfig()
	secretKey = []byte(cfg.SecretKey.Secret)
)

func (s *service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		log.Printf("Error decoding login request: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var hashedPassword string
	err = s.db.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", loginReq.Username).Scan(&hashedPassword)
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
	err = s.db.QueryRow(context.Background(), "SELECT id, role FROM users WHERE username = $1", loginReq.Username).Scan(&user.ID, &user.Role)
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
