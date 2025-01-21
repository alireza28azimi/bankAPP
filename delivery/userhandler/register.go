package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"main.go/dto"
)

func (h Handler) userRegister(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "page not found",
		})

	}
	return nil
}
