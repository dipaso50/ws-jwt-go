package infraestructure

import (
	"jwtMiddleware/jwt/model"
	"testing"
)

const secret = "hola"

func Test_shouldBeValid(t *testing.T) {
	generator := NewTokenGenerator(secret)
	validator := NewTokenValidatorService(secret)

	username := "jhon"

	tok, err := generator.Generate(username)

	if err != nil {
		t.Errorf("Error creando token %v", err)
	}

	_, err = validator.Validate(tok)

	if err != nil {
		t.Errorf("Error validando token %v", err)
	}
}

func Test_souldBeInvalid(t *testing.T) {

	validator := NewTokenValidatorService(secret)

	_, err := validator.Validate(model.Token{"kdkdk"})

	if err == nil {
		t.Errorf("Expecting error on validate %v", err)
	}
}
