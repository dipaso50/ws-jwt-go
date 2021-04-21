package model

type TokenGenerator interface {
	Generate(name string) (Token, error)
}
