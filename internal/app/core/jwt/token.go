package jwt

import "FaisalBudiono/go-jwt/internal/app/domain"

func NewTokenGen() *token {
	return &token{}
}

type token struct{}

func (t *token) Gen(u domain.User) (domain.Token, error) {
	panic("unimplemented")
}
