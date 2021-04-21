package handler

import (
	"context"
	"jwtMiddleware/jwt/model"
	"log"
	"net/http"
	"strings"
)

type AuthTokenMiddleware struct{}

func TokenValidationMiddleware(ctx context.Context, tokenValidator model.TokenValidator, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			log.Println("Missing Authorization Header")
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := tokenValidator.Validate(model.Token{AccessToken: tokenString})

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		newCtx := context.WithValue(ctx, currentClaims, claims)
		req := r.WithContext(newCtx)
		next.ServeHTTP(w, req)
	})
}
