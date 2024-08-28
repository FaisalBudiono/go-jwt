package httplib

import (
	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(code int, err error) error {
	return echo.NewHTTPError(code, errorResponse{
		Message: err.Error(),
	})
}
