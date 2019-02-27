package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type errorResponse struct {
	Error string `json:"error"`
}

func TestGetOk(t *testing.T) {
	uri := "/test-controllers/get-ok"
	router := route("GET", uri, GetOk)

	testcases := []struct {
		desc     string
		respStr  string
		respCode int
	}{
		{"Test Ok - 1", "ok", http.StatusOK},
		{"Test Ok - 2", "ok", http.StatusOK},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			req, _ := http.NewRequest("GET", uri, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)

			respBody, _ := ioutil.ReadAll(w.Result().Body)
			respStr := string(respBody)
			assert.Equal(t, tc.respStr, respStr)
		})
	}
}

func TestFavIcon(t *testing.T) {
	uri := "/test-controllers/favicon.ico"
	router := route("GET", uri, FavIcon)

	testcases := []struct {
		desc            string
		respContentType string
		respCode        int
	}{
		{"Test Fav Icon - 1", "image/x-icon", http.StatusOK},
		{"Test Fav Icon - 2", "image/x-icon", http.StatusOK},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			req, _ := http.NewRequest("GET", uri, nil)
			req.Header.Set("Content-type", "WimTest8()")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)

			resp := w.Result()
			respBody, _ := ioutil.ReadAll(resp.Body)
			assert.NotZero(t, len(respBody))

			assert.Equal(t, tc.respContentType, resp.Header.Get("Content-Type"))
		})
	}
}

func route(method, uri string, handler gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	if method == "POST" {
		router.POST(uri, handler)
	} else {
		router.GET(uri, handler)
	}
	return router
}
