package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricRouterInit(e *gin.Engine) {
	// 暴露 `/metrics` 供 Prometheus 采集
	e.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
