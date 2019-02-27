package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/controllers"
)

// MakeCommonAPI -
func MakeCommonAPI(r *gin.Engine) {
	r.GET("/ok", controllers.GetOk)
	r.GET("/favicon.ico", controllers.FavIcon)
}
