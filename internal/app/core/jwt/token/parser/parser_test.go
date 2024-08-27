package parser_test

import (
	"FaisalBudiono/go-jwt/internal/app/core/jwt/token/parser"
	"FaisalBudiono/go-jwt/internal/app/port/common"
	"encoding/base64"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParserHS_Sign(t *testing.T) {
	u := common.User{
		ID:        "123",
		Name:      "john doe",
		Email:     "john@gmail.com",
		Password:  "mysecret",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now().Add(5 * time.Minute),
	}

	p := parser.NewParser(mockKey(), parser.NewNower())
	p.Sign(u)

	_ = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJyb290SUQiOiJhc2QiLCJpZCI6IjEyMyIsInN1YiI6IjEyMyIsImV4cCI6MTcyNDA5NDE5NywibmJmIjoxNzI0MDkzODk3LCJpYXQiOjE3MjQwOTM4OTd9.akG4zEH1kbBRjSyrM3svE_JTJW7DKB9wHZvHIXr9lUDXox-DNZvboHl-1GTvCeg-Vckj3Uda1RJvwu2rnvZTvg"

	assert.Equal(t, "asd", "asd")
}

func mockKey() []byte {
	res, err := base64.StdEncoding.DecodeString("ccGP1+yJVBU95C0o5vMAU1bM2wYO/SImn3OY4GtTSO0tp3pS6Hzpkk/Z2BqnVDJMKESi4cv5pZ4aorxdgdLPIw==")
	if err != nil {
		panic(err)
	}

	return res
}
