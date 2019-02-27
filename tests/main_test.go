package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wmh/my-gin-example/app/routes"
)

func TestFlow(t *testing.T) {
	r := gin.Default()
	routes.MakeCommonAPI(r)
	routes.MakeExampleAPI(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ok", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/example/auth/example", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/example/auth/example", nil)
	req.Header.Add("X-Auth", "PASS")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
