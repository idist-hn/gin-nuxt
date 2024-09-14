package routes

import (
	"github.com/gin-gonic/gin"
	"idist-core/app/controllers/cie"
)

func CommonRoutes(router *gin.RouterGroup) {
	router.GET("/sub-menu", cie.GetSubMenu)
}
