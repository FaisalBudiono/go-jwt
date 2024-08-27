package port

import "context"

type LoginInput struct {
	Ctx      context.Context
	Email    string
	Password string
}

type RegisterInput struct {
	Ctx      context.Context
	Name     string
	Email    string
	Password string
}
