package token_test

import (
	"FaisalBudiono/go-jwt/internal/app/core/jwt/token"
	"FaisalBudiono/go-jwt/internal/app/port/common"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	first := "random"
	uuidGen := newIDGen(first)
	u := common.User{
		ID:        "asd",
		Name:      "zxc",
		Email:     "asd@asd.com",
		Password:  "asd",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	signer := NewSigner(t, u)

	gen := token.NewGen(uuidGen, signer)

	res, err := gen.Create(u)

	expectedAccessToken, _ := signer.Sign(u)
	expectedToken := token.Token{
		AccessToken: expectedAccessToken,
		RefreshToken: token.RefreshToken{
			ID:       first,
			RootID:   first,
			ParentID: first,
		},
	}

	assert.Equal(t, expectedToken, res)
	assert.Nil(t, err)
}

type signer struct {
	tt *testing.T
	u  common.User
}

func (s *signer) Sign(u common.User) (string, error) {
	assert.Equal(s.tt, u, s.u)

	return u.ID + u.Email, nil
}

func NewSigner(
	tt *testing.T,
	u common.User,
) *signer {
	return &signer{
		tt: tt,
		u:  u,
	}
}

type idGenerator struct {
	stacks []string
}

func newIDGen(stacks ...string) *idGenerator {
	return &idGenerator{
		stacks: stacks,
	}
}

func (g *idGenerator) UUID() (string, error) {
	if len(g.stacks) == 0 {
		return "", errors.New("empty stack")
	}

	stack := g.stacks[0]
	g.stacks = g.stacks[1:]

	return stack, nil
}
