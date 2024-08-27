package main

import (
	"FaisalBudiono/go-jwt/internal/env"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func main() {
	env.Bind()

	sec64 := viper.GetString("JWT_SECRET_BASE64")
	secret, err := base64.StdEncoding.DecodeString(sec64)
	if err != nil {
		fmt.Printf("Kambing err %#v\n", err)
	}

	iat := time.Now()
	exp := time.Now().Add(time.Minute * 5)

	claim := &Claims{
		ID: "1",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "ucul",
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	fmt.Printf("Kambing claims %#v\n", t.Claims)

	res, err := t.SignedString(secret)
	fmt.Printf("Kambing res %#v\n", res)
	fmt.Printf("Kambing err %#v\n\n\n", err)

	tokenString := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJzdWIiOiJ1Y3VsIiwiZXhwIjoxNzI0MDkyMDIxLCJpYXQiOjE3MjQwOTE3MjF9.poETPC99WMELrpzHQIeV-N50pf3SNj2EGEDpjdUXgQGs_TDrkcAPHNbjyzAowcNyt4IQcgUWY5sCsGl5E4BOfg"

	rawClaim, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		fmt.Printf("Kambing error claim %#v\n", err)
	}

	resClaim, ok := rawClaim.Claims.(*Claims)
	if !ok {
		fmt.Printf("gagal parsing %#v\n", err)
	}

	fmt.Printf("HUAHAHA claim res %#v\n", resClaim)
}
