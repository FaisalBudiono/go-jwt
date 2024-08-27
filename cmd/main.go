package main

import (
	"FaisalBudiono/go-jwt/internal/app/adapter"
	"FaisalBudiono/go-jwt/internal/app/core"
	"FaisalBudiono/go-jwt/internal/app/core/pwhash/argon"
	"FaisalBudiono/go-jwt/internal/app/port"
	"FaisalBudiono/go-jwt/internal/db"
	"FaisalBudiono/go-jwt/internal/env"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	env.Bind()
	db.Init()
}

type ErrorRes struct {
	Error string `json:"error"`
}

func main() {
	e := echo.New()

	adb := adapter.NewDB(db.DB)
	hasher := argon.New()
	cauth := core.NewAuth(adb, hasher)

	e.Use(middleware.Logger())

	e.POST("/sign-up", func(c echo.Context) error {
		type req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var r req
		err := c.Bind(&r)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorRes{err.Error()})
		}

		u, err := cauth.Reg(port.RegisterInput{
			Ctx:      c.Request().Context(),
			Name:     r.Name,
			Email:    r.Email,
			Password: r.Password,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorRes{err.Error()})
		}

		return c.JSON(http.StatusOK, struct {
			ID        string    `json:"id"`
			Name      string    `json:"name"`
			Email     string    `json:"email"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
		}{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
