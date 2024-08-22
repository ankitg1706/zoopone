package main

import (
	"github.com/ankitg1706/zoopone/api"
	"github.com/ankitg1706/zoopone/controller"
	"github.com/gin-gonic/gin"
)

// @title Managment
// @version 1.0
// @description API for managing School operations
// @host localhost:8000
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token
func main() {
	api := api.APIRoutes{}
	controller := controller.Server{}
	routes := gin.Default()
	api.StartApp(routes, controller)

	routes.Run(":8000")
	// fmt.Printf("main server = %v\n", api)
}
