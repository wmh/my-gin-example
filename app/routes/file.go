package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/controllers"
	"github.com/wmh/my-gin-example/app/services"
)

func MakeFileAPI(r *gin.Engine) {
	files := r.Group("/v2/files")
	files.Use(services.JWTAuthMiddleware())
	{
		files.POST("/upload", controllers.UploadFile)
		files.POST("/upload/multiple", controllers.UploadMultipleFiles)
		files.GET("", controllers.ListFiles)
		files.GET("/:id", controllers.GetFile)
		files.GET("/:id/download", controllers.DownloadFile)
		files.DELETE("/:id", controllers.DeleteFile)
	}
}
