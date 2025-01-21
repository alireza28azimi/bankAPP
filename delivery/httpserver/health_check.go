package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func healthCheck(c echo.Context) {
	c.JSON(http.StatusOK, echo.Map{
		"message": "every thing is good",
	})
}
