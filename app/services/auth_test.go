package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExampleAuthSuccess(t *testing.T) {
	uri := "/test-example-auth"
	resp := "ok"
	requestCnt := 10

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(ExampleAuth())
	router.GET(uri, func(c *gin.Context) {
		c.String(http.StatusOK, resp)
	})

	for i := 0; i < requestCnt; i++ {
		t.Run(fmt.Sprintf("Parallel Should Succeed Example Auth Request: %d", i), func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", uri, nil)
			req.Header.Add("X-Auth", "PASS")
			router.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, resp, w.Body.String())
		})
	}
}

func TestExampleAuthFailed(t *testing.T) {
	uri := "/test-example-auth-failed"
	resp := ""
	requestCnt := 10

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(ExampleAuth())
	router.GET(uri, func(c *gin.Context) {
		c.String(http.StatusOK, resp)
	})

	for i := 0; i < requestCnt; i++ {
		t.Run(fmt.Sprintf("Parallel Should Succeed Example Auth Request: %d", i), func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", uri, nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, http.StatusUnauthorized, w.Code)
			assert.Equal(t, resp, w.Body.String())
		})
	}
}
