package userhandler

import "github.com/gin-gonic/gin"

func (h Handler) SetUserRoute(router *gin.RouterGroup) {
	userG := router.Group("/users")
	userG.POST("/register", h.userRegister)
	userG.POST("/login", h.userLogin)
	userG.GET("profile", h.userProfile)
}
