package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/controllers"
	"github.com/wmh/my-gin-example/app/services"
)

// MakeExampleAPI -
func MakeExampleAPI(r *gin.Engine) {
	v1 := r.Group("/v1/example")
	{
		v1.GET("/hello", controllers.Hello)
		v1.POST("/hello", controllers.PostHello)
		v1.GET("/longRequest", controllers.LongRequest)

		exampleAuth := v1.Group("/auth/example")
		exampleAuth.Use(services.ExampleAuth())
		{
			exampleAuth.GET("", controllers.AuthHello)
		}
	}
}
