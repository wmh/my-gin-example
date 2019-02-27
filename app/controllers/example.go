package controllers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/core"
)

// Hello -
func Hello(c *gin.Context) {
	msg := core.H{"ts": time.Now().Unix(), "path": c.Request.RequestURI}
	core.Log("getHello", msg)
	c.JSON(http.StatusOK, gin.H{"msg": "hello world"})
}

// PostHello -
func PostHello(c *gin.Context) {
	msg := core.H{"ts": time.Now().Unix(), "path": c.Request.RequestURI}
	core.Log("postHello", msg)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// LongRequest -
func LongRequest(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// AuthHello -
func AuthHello(c *gin.Context) {
	uri := c.Request.RequestURI
	var body []byte
	if c.Request.Body != nil {
		body, _ = ioutil.ReadAll(c.Request.Body)
	}

	msg := core.H{"ts": time.Now().Unix(), "path": uri}
	core.Log("authHello", msg)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "login success",
		"uri":  uri,
		"body": string(body),
	})
}
