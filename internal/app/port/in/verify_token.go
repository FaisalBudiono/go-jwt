package in

import "context"

type VerifyTokenPort interface {
	Ctx() (context.Context, error)
	AccessToken() (string, error)
}
