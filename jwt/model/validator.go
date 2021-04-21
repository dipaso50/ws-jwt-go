package model

type TokenValidator interface {
	Validate(tok Token) (MyClaims, error)
}
