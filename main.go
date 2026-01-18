package main

import (
	"exchangeapp/config"
	// "fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	// fmt.Println(config.AppConfig.App.Port)
	 r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	c.JSON(200, gin.H{
	"message": "pong",
	})
	})
	r.Run()
}
