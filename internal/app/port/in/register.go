package in

import (
	"context"
)

type RegisterPort interface {
	Ctx() (context.Context, error)
	Name() (string, error)
	Email() (string, error)
	Password() (string, error)
}
