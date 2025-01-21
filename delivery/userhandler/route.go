package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetUserRoute(e *echo.Echo) {
	userG := e.Group("/users")
	userG.POST("/register", h.userRegister)
	userG.POST("/login", h.userLogin)

}
