package middleware

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
