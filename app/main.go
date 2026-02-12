package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"ilicense-lite/bootstrap"
	"ilicense-lite/config"
	"ilicense-lite/middleware"
	"ilicense-lite/router"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// @title           Gin App API
// @version         1.0
// @description     这是gin-app服务的接口文档.
// @host            localhost:8080
// @BasePath        /api
func main() {
	file := flag.String("config", "etc/app.yaml", "path to the configuration file")
	flag.Parse()
	bootstrap.Init(*file)
	r := gin.New()
	r.Use(otelgin.Middleware("ilicense-lite"))
	r.Use(middleware.MetricsMiddleware())
	r.Use(middleware.PrometheusMiddleware())
	r.Use(gin.Logger())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(cors.Default())
	r.Use(middleware.AuthMiddleware())
	router.Init(r)
	srv := endless.NewServer(fmt.Sprintf(":%d", config.Config.App.Port), r)
	srv.RegisterOnShutdown(func() {
		fmt.Println("shutting down server...")
	})
	srv.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGINT, func() {
		log.Println("model ctrl+c before shutting down server...")
	})
	srv.RegisterSignalHook(endless.POST_SIGNAL, syscall.SIGINT, func() {
		log.Println("model ctrl+c after shutting down server...")
	})
	srv.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGTERM, func() {
		log.Println("model kill before shutting down server...")
	})
	srv.RegisterSignalHook(endless.POST_SIGNAL, syscall.SIGTERM, func() {
		log.Println("model kill after shutting down server...")
	})
	srv.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGUSR1, func() {
		log.Println("model usr-1 before shutting down server...")
	})
	srv.RegisterSignalHook(endless.POST_SIGNAL, syscall.SIGUSR1, func() {
		log.Println("model usr-1 after shutting down server...")
	})
	srv.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGUSR2, func() {
		log.Println("model usr-2 before shutting down server...")
	})
	srv.RegisterSignalHook(endless.POST_SIGNAL, syscall.SIGUSR2, func() {
		log.Println("model usr-2 after shutting down server...")
	})
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server stopping about %v", err)
	}
}
