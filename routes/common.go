package routes

import (
	"github.com/gin-gonic/gin"
	"idist-core/app/controllers/admin"
	"idist-core/middleware"
)

func CommonRoutes(router *gin.RouterGroup) {
	router.GET("/provinces", middleware.Gate("", ""), admin.ListProvinces)
	router.GET("/provinces/id", middleware.Gate("", ""), admin.ReadProvince)
	router.PUT("/provinces/id", middleware.Gate("", ""), admin.UpdateProvince)

	router.POST("/tuyen-sinh", middleware.Gate("", ""), admin.CreateTuyenSinh)
}
