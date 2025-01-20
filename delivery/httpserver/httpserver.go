package httpserver

import (
	"github.com/gin-gonic/gin"
	"main.go/config"
)

type Server struct {
	confg       config.Config
	userhandler UserHandler
}

func (s Server) Serve() {
	// Create a new Gin router
	router := gin.Default()

	// Define a simple GET route
	router.GET("/heathcheck", healthCheck)

	s.userhandler.SetUserRoute(router)

	// Start the server on port 8080
	router.Run(":8080")
}
