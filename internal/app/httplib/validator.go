package httplib

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type customValidator struct {
	validator *validator.Validate
}

func NewValidator() *customValidator {
	v := validator.New(validator.WithRequiredStructEnabled())

	return &customValidator{
		validator: v,
	}
}

func (v *customValidator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return nil
}
