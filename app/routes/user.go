package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/controllers"
	"github.com/wmh/my-gin-example/app/services"
)

func MakeUserAPI(r *gin.Engine) {
	auth := r.Group("/v2/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	users := r.Group("/v2/users")
	users.Use(services.JWTAuthMiddleware())
	{
		users.GET("/profile", controllers.GetProfile)
		users.PUT("/profile", controllers.UpdateProfile)
		users.GET("", services.RequireRole("admin"), controllers.ListUsers)
	}
}
