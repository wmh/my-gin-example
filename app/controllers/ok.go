package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

const faviconData = "AAABAAEAEBAQAAAAAAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAA////AJhZOwDMeFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIiIiIiIiIAIzETMzETMgAjMRMzMRMyACIhEiIhEiIAIiESIiESIgAiIREiIRIiACIhERIhEiIAIiESERESIgAiIRIhESIiACIhEiIiIiIAIiESIiIiIgAiIRIiIiIiACIiIiIiIiIAIiIiIiIiIgAAAAAAAAAAD//wAAgAEAAIABAACAAQAAgAEAAIABAACAAQAAgAEAAIABAACAAQAAgAEAAIABAACAAQAAgAEAAIABAAD//wAA"

var favicon []byte
var faviconLength int64

func init() {
	favicon, _ = base64.StdEncoding.DecodeString(faviconData)
	faviconLength = int64(len(favicon))
}

// GetOk - response ok
func GetOk(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// FavIcon - FavIcon for browser
func FavIcon(c *gin.Context) {
	c.Header("Cache-Control", "max-age=315360000")
	c.Data(http.StatusOK, "image/x-icon", favicon)
}
