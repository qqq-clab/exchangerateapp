package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"Message": "login successsful",
			})
		})

		auth.POST("/register", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"Message": "register successsful",
			})
		})
	}

	return router
}
