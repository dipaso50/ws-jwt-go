package infraestructure

import (
	"jwtMiddleware/jwt/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenGeneratorService struct {
	secret string
}

func NewTokenGenerator(secret string) TokenGeneratorService {
	return TokenGeneratorService{secret: secret}
}

func (tgs TokenGeneratorService) Generate(name string) (model.Token, error) {

	mclaims := jwt.MapClaims{}
	mclaims["name"] = name
	mclaims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, mclaims)
	token, err := at.SignedString([]byte(tgs.secret))

	if err != nil {
		return model.Token{}, err
	}

	return model.Token{AccessToken: token}, nil
}
