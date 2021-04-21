package infraestructure

import (
	"fmt"
	"jwtMiddleware/jwt/model"

	"github.com/dgrijalva/jwt-go"
)

type TokenValidatorService struct {
	secret string
}

func NewTokenValidatorService(sec string) TokenValidatorService {
	return TokenValidatorService{secret: sec}
}

func (tvs TokenValidatorService) Validate(tok model.Token) (model.MyClaims, error) {
	claims := jwt.MapClaims{}

	var jwtSecretKey = []byte(tvs.secret)

	token, err := jwt.ParseWithClaims(tok.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return model.MyClaims{}, err
	}

	if token != nil {
		return model.MyClaims{
			Name: claims["name"].(string),
		}, nil
	}

	return model.MyClaims{}, fmt.Errorf("Invalid token")
}
