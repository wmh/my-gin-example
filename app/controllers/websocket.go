package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/wmh/my-gin-example/app/core"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSMessage struct {
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		core.ErrorLog("websocket", err.Error())
		return
	}
	defer conn.Close()

	clientID := c.ClientIP()
	core.Log("websocket", core.H{"action": "connected", "client": clientID})

	welcomeMsg := WSMessage{
		Type:      "welcome",
		Data:      map[string]string{"message": "Connected to WebSocket server"},
		Timestamp: time.Now().Unix(),
	}
	if err := conn.WriteJSON(welcomeMsg); err != nil {
		core.ErrorLog("websocket", err.Error())
		return
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			var msg WSMessage
			err := conn.ReadJSON(&msg)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					core.ErrorLog("websocket", err.Error())
				}
				return
			}

			core.Log("websocket", core.H{"action": "received", "client": clientID, "message": msg})

			response := WSMessage{
				Type: "echo",
				Data: map[string]interface{}{
					"received": msg,
					"echo":     "Message received",
				},
				Timestamp: time.Now().Unix(),
			}

			if err := conn.WriteJSON(response); err != nil {
				core.ErrorLog("websocket", err.Error())
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			core.Log("websocket", core.H{"action": "disconnected", "client": clientID})
			return
		case <-ticker.C:
			msg := WSMessage{
				Type: "ping",
				Data: map[string]interface{}{
					"server_time": time.Now().Format(time.RFC3339),
					"uptime":      "active",
				},
				Timestamp: time.Now().Unix(),
			}
			if err := conn.WriteJSON(msg); err != nil {
				core.ErrorLog("websocket", err.Error())
				return
			}
		}
	}
}

func BroadcastExample(c *gin.Context) {
	var message map[string]interface{}
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	core.Log("broadcast", message)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Broadcast initiated (in production, this would send to all connected clients)",
		"data":    message,
	})
}
