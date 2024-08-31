package cacher

import (
	"FaisalBudiono/go-jwt/internal/app/domain"
	"context"
)

type redis struct {
}

func (r *redis) Cache(ctx context.Context, t domain.Token) error {
	panic("unimplemented")
}

func NewRedis() *redis {
	return &redis{}
}
