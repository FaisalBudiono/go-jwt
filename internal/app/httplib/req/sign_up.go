package req

import "context"

type SignUp struct {
	Context   context.Context `json:"-"`
	FName     string          `json:"name" validate:"required"`
	FEmail    string          `json:"email" validate:"required"`
	FPassword string          `json:"password" validate:"required"`
}

func (s *SignUp) Ctx() (context.Context, error) {
	return s.Context, nil
}

func (s *SignUp) Email() (string, error) {
	return s.FEmail, nil
}

func (s *SignUp) Name() (string, error) {
	return s.FName, nil
}

func (s *SignUp) Password() (string, error) {
	return s.FPassword, nil
}
