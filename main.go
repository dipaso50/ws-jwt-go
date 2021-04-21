package main

import (
	"context"
	"jwtMiddleware/handler"
	"jwtMiddleware/jwt/infraestructure"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	secret = os.Getenv("TOKEN_TEST_SECRET")
	port   = os.Getenv("TOKEN_TEST_PORT")
)

func main() {
	r := mux.NewRouter()

	ctx := context.Background()

	tokenGenerator := infraestructure.NewTokenGenerator(secret)
	tokenValidator := infraestructure.NewTokenValidatorService(secret)

	r.Handle("/token", handler.CreateToken(tokenGenerator)).Methods("GET")
	r.Handle("/sayHello", handler.TokenValidationMiddleware(ctx, tokenValidator, handler.SayHello())).Methods("GET")

	log.Println("Listening in port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
