package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/core"
)

// ExampleAuth An example shows how auth middleware work
func ExampleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if auth := c.Request.Header.Get("X-Auth"); auth != "PASS" {
			core.Log("auth", "Example auth failed!")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
