package handler

import (
	"encoding/json"
	"jwtMiddleware/jwt/model"
	"log"
	"net/http"
)

type User struct {
	Name *string `json:"name"`
}

type ContextKey string

const currentClaims ContextKey = "currentClaims"

func CreateToken(tokenGenerator model.TokenGenerator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		token, err := tokenGenerator.Generate(*user.Name)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(&token)
	})
}
