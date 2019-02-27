package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/core"
	"github.com/wmh/my-gin-example/app/routes"
)

func main() {
	if core.ConfBool("logs.disable_default_writer") {
		gin.DefaultWriter = ioutil.Discard
	}

	r := gin.Default()
	routes.MakeCommonAPI(r)
	routes.MakeExampleAPI(r)

	appPort := core.ConfString("APP_PORT")
	srv := &http.Server{
		Addr:    ":" + appPort,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			core.ErrorLog("main", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
