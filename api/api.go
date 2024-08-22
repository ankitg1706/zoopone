package api

import (
	"github.com/ankitg1706/zoopone/controller"
	_ "github.com/ankitg1706/zoopone/docs"
	"github.com/ankitg1706/zoopone/store"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type APIRoutes struct {
	Server controller.ServerOperations
}

func (api *APIRoutes) StartApp(router *gin.Engine, server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user routs
	api.UserRouts(router)

}
