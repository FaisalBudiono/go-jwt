package main

import (
	"FaisalBudiono/go-jwt/internal/db"
	"FaisalBudiono/go-jwt/internal/env"
)

func init() {
	env.Bind()
	db.Init()
}

func main() {
}
