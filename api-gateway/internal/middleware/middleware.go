package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware(cfg *config.Config, next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			log.Printf("JWT не найден: %s", err)
			http.Error(w, "JWT не найден", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.SecretKey.Secret), nil
		})
		if err != nil {
			log.Printf("JWT: %s\n", cfg.SecretKey.Secret)
			log.Printf("JWT невалиден: %s", err)
			http.Error(w, "JWT невалиден", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {

			log.Printf("JWT невалиден: %s", err)
			http.Error(w, "JWT невалиден", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims.UserID)
		next(w, r.WithContext(ctx))
	}
}
