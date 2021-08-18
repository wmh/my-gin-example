package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/controllers"
	"github.com/wmh/my-gin-example/app/services"
)

// MakeExampleAPI -
func MakeExampleAPI(r *gin.Engine) {
	v1 := r.Group("/v1/example")
	{
		v1.GET("/hello", controllers.Hello)
		v1.POST("/hello", controllers.PostHello)
		v1.GET("/longRequest", controllers.LongRequest)

		exampleAuth := v1.Group("/auth/example")
		exampleAuth.Use(services.ExampleAuth())
		{
			exampleAuth.GET("", controllers.AuthHello)
		}
	}

	r.Any("/proxypass/*proxyPath", proxy)
}

func proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:8010")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
