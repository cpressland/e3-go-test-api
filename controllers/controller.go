package controllers

import (
	"go-test-api/repositories"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	api            *gin.Engine
	userRepository repositories.UserRepository
}

func NewController(api *gin.Engine, userRepository repositories.UserRepository) Controller {
	controller := &controller{
		api:            api,
		userRepository: userRepository,
	}

	api.POST("/users", controller.Create)
	api.GET("/users", controller.List)
	api.GET("/users/:id", controller.Get)
	api.PUT("/users/:id", controller.Update)
	api.DELETE("/users/:id", controller.Delete)

	return controller
}
