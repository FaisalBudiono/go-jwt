package token

import "FaisalBudiono/go-jwt/internal/app/port/common"

type (
	RefreshToken struct {
		ID       string
		RootID   string
		ParentID string
	}

	Token struct {
		AccessToken  string
		RefreshToken RefreshToken
	}
)

type idGenerator interface {
	UUID() (string, error)
}

type jwtSigner interface {
	Sign(u common.User) (string, error)
}

type gen struct {
	idgen     idGenerator
	jwtSigner jwtSigner
}

func NewGen(idgen idGenerator, jwtSigner jwtSigner) *gen {
	return &gen{
		idgen:     idgen,
		jwtSigner: jwtSigner,
	}
}

func (g *gen) Create(u common.User) (Token, error) {
	accessToken, err := g.jwtSigner.Sign(u)
	if err != nil {
		return Token{}, err
	}

	uuid, err := g.idgen.UUID()
	if err != nil {
		return Token{}, err
	}

	refreshToken := RefreshToken{
		ID:       uuid,
		RootID:   uuid,
		ParentID: uuid,
	}

	return Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
