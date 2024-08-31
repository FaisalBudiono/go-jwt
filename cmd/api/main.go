package main

import (
	"FaisalBudiono/go-jwt/internal/app/adapter/cacher"
	"FaisalBudiono/go-jwt/internal/app/adapter/pqrepo"
	"FaisalBudiono/go-jwt/internal/app/core"
	"FaisalBudiono/go-jwt/internal/app/core/hasher/argon"
	"FaisalBudiono/go-jwt/internal/app/core/jwt"
	"FaisalBudiono/go-jwt/internal/app/httplib"
	"FaisalBudiono/go-jwt/internal/app/httplib/req"
	"FaisalBudiono/go-jwt/internal/app/httplib/reso"
	"FaisalBudiono/go-jwt/internal/db"
	"FaisalBudiono/go-jwt/internal/env"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	env.Bind()
}

func main() {
	e := echo.New()
	e.Validator = httplib.NewValidator()

	dbConn := db.PostgresConn()
	jwtSigner := jwt.NewTokenGen()
	redisClient := cacher.NewRedis()

	auth := core.NewAuth(
		dbConn,
		argon.New(),
		pqrepo.New(dbConn),
		jwtSigner,
		redisClient,
	)

	e.Use(middleware.Logger())

	e.POST("/sign-up", func(c echo.Context) error {
		r := &req.SignUp{
			Context: c.Request().Context(),
		}

		err := c.Bind(r)
		if err != nil {
			return httplib.NewErrorResponse(http.StatusUnprocessableEntity, err)
		}

		err = c.Validate(r)
		if err != nil {
			return err
		}

		res, err := auth.Reg(r)
		if err != nil {
			return httplib.NewErrorResponse(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, reso.MapUser(res))
	})

	e.Logger.Fatal(e.Start(":" + viper.GetString("APP_PORT")))
}
