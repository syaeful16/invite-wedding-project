package middlewares

import (
	"context"
	"invite-wed/configs"
	"invite-wed/helpers"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const (
	IdKey       ContextKey = "id"
	UsernameKey ContextKey = "username"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helpers.JsonResponse(w, http.StatusUnauthorized, response)
				return
			}
		}

		// get token from cookie
		tokenString := c.Value
		claims := &configs.JWTClaims{}

		// parsing jwt token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return configs.JWT_KEY, nil
		})

		if err != nil {
			response := map[string]string{"message": "Unauthorized"}
			helpers.JsonResponse(w, http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helpers.JsonResponse(w, http.StatusUnauthorized, response)
			return
		}
		ctx := context.WithValue(r.Context(), IdKey, claims.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
