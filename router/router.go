package router

import (
	"exchangeapp/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)

		auth.POST("/register", controllers.Register)
	}

	return router
}
