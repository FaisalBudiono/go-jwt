package parser

import (
	"FaisalBudiono/go-jwt/internal/app/port/common"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type Parser struct {
	secretKey []byte
	nower     nower
}

type nower interface {
	Now() time.Time
}

type nowerReal struct{}

func NewNower() *nowerReal {
	return &nowerReal{}
}

func (n *nowerReal) Now() time.Time {
	return time.Now()
}

func NewParser(secretKey []byte, nower nower) *Parser {
	time.Now()
	return &Parser{
		secretKey: secretKey,
		nower:     nower,
	}
}

func (p *Parser) Sign(u common.User) (string, error) {
	now := p.nower.Now()

	claim := &Claims{
		ID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.ID,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 5)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	res, err := t.SignedString(p.secretKey)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (p *Parser) Decode(token string) (Claims, error) {
	return Claims{}, nil
}
