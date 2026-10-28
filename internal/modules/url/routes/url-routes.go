package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SubhamSaha99/url-shortner/internal/modules/url/controllers"
)

func RegisterURLRoutes(
	router *gin.Engine,
	controller *controllers.URLController,
) {
	api := router.Group("/api/v1")

	api.POST("/urls", controller.Create)

	router.GET("/:code", controller.Redirect)

	router.POST("/delete", controller.DeleteURL)
}