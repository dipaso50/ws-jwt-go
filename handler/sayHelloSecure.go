package handler

import (
	"jwtMiddleware/jwt/model"
	"net/http"
)

func SayHello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentClaims := r.Context().Value(currentClaims).(model.MyClaims)
		w.Write([]byte("Hello " + currentClaims.Name))
	})
}
