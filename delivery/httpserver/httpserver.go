package httpserver

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main.go/config"
	"main.go/delivery/userhandler"
	"main.go/service/authservice"
	"main.go/service/userservice"
	"main.go/service/uservalidator"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func new() {

}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator) Server {
	return Server{
		config:      config,
		userHandler: userhandler.New(authSvc, userSvc, userValidator, config.Auth),
	}
}
func (s Server) Serve() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health_check", healthCheck)

	s.userHandler.SetUserRoute(e)

	slog.Info("Attempting to start server on port 2002...")
	// Start server
	if err := e.Start(":2002"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
