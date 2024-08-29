package routerProvider

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"idist-core/app/providers/loggerProvider"
	_ "idist-core/docs"
	middlewares2 "idist-core/middleware"
	"idist-core/routes"
)

func Init(router *gin.Engine) {
	loggerProvider.GetLogger().Info("------------------------------------------------------------")
	router.Use(requestid.New())
	router.Use(middlewares2.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares2.CorsMiddleware())
	router.Use(middlewares2.ConfigsMiddleware())
	//docs.SwaggerInfo.BasePath = "swagger/v1"
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api := router.Group("api/v1")
	routes.AdminRoutes(api.Group("admin"))
	routes.AuthRoutes(api.Group("auth"))
	routes.CommonRoutes(api.Group("common"))
	loggerProvider.GetLogger().Info("------------------------------------------------------------")

}
