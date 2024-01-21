package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitGinEngine() *gin.Engine {
	app := gin.New()

	app.MaxMultipartMemory = 8 << 20
	app.NoMethod(NoMethodHandler())
	app.NoRoute(NoRouteHandler())

	app.Use(gin.Recovery())
	app.Use(cors.Default())

	// Swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
