package routes

import (
	"github.com/gin-gonic/gin"
	"go-assignment/handlers"
)


func RegisterRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", handler.CreateUser)
		userRoutes.GET("/", handler.ListUsers)
        userRoutes.GET("/:id", handler.GetUser)
        userRoutes.PUT("/", handler.UpdateUser)
        userRoutes.DELETE("/:id", handler.DeleteUser)

	}
}