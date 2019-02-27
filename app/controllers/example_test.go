package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOk1(t *testing.T) {
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
