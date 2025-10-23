package routes

import (
	"gin-gorm-swagger-api/controllers"
	"gin-gorm-swagger-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	users.Use(middlewares.JWTAuthMiddleware())
	{
		users.GET("", controllers.GetUsers)
		users.GET("/:id", controllers.GetUserByID)
		users.POST("", controllers.CreateUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
