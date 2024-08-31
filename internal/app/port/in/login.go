package in

import (
	"context"
)

type LoginPort interface {
	Ctx() (context.Context, error)
	Email() (string, error)
	Password() (string, error)
}
