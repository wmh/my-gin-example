package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/core"
	"github.com/wmh/my-gin-example/app/models"
	"github.com/wmh/my-gin-example/app/routes"
	"github.com/wmh/my-gin-example/app/services"
)

func main() {
	if core.ConfBool("logs.disable_default_writer") {
		gin.DefaultWriter = io.Discard
	}

	if err := core.InitDB(); err != nil {
		core.ErrorLog("main", "Failed to initialize database: "+err.Error())
		os.Exit(1)
	}
	defer core.CloseDB()

	if err := core.DB.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		core.ErrorLog("main", "Failed to migrate database: "+err.Error())
		os.Exit(1)
	}

	services.InitRateLimiter(100, time.Minute)

	r := gin.Default()
	r.Use(services.RateLimitMiddleware())

	routes.MakeCommonAPI(r)
	routes.MakeExampleAPI(r)
	routes.MakeUserAPI(r)
	routes.MakeProductAPI(r)
	routes.MakeWebSocketAPI(r)

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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	core.Log("main", "Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		core.ErrorLog("main", "Server forced to shutdown: "+err.Error())
	}

	core.Log("main", "Server exited")
}
