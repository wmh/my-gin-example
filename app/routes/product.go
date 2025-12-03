package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/controllers"
	"github.com/wmh/my-gin-example/app/services"
)

func MakeProductAPI(r *gin.Engine) {
	products := r.Group("/v2/products")
	{
		products.GET("", controllers.ListProducts)
		products.GET("/:id", controllers.GetProduct)

		protected := products.Group("")
		protected.Use(services.JWTAuthMiddleware())
		{
			protected.POST("", controllers.CreateProduct)
			protected.PUT("/:id", controllers.UpdateProduct)
			protected.DELETE("/:id", services.RequireRole("admin"), controllers.DeleteProduct)
		}
	}
}
