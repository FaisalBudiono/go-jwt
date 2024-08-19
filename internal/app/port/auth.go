package port

import "context"

type RegisterInput struct {
	Ctx      context.Context
	Name     string
	Email    string
	Password string
}
