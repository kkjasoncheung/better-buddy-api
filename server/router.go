package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kkjasoncheung/better-buddy-api/controllers"
)

// NewRouter returns a new router with middleware and endpoints.
func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userController := new(controllers.UserController)
			userGroup.GET("/", userController.Retrieve)
			userGroup.GET("/:id", userController.RetrieveByID)
			userGroup.POST("/", userController.Create)
			userGroup.PATCH("/:id", userController.Update)
			userGroup.DELETE("/:id", userController.Delete)
		}
	}

	return router
}
