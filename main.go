package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
	// "fmt"
)

func main() {
	config.InitConfig()
	// fmt.Println(config.AppConfig.App.Port)
	r := router.SetupRouter()
	r.Run()
}
