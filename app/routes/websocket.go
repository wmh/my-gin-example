package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/controllers"
)

func MakeWebSocketAPI(r *gin.Engine) {
	ws := r.Group("/ws")
	{
		ws.GET("/connect", controllers.WebSocketHandler)
		ws.POST("/broadcast", controllers.BroadcastExample)
	}
}
