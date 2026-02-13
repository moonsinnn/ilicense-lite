package router

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	DocsRouterInit(r)
	MetricRouterInit(r)
	UserRouterInit(r)
	ProductRouterInit(r)
	CustomerRouterInit(r)
	LicenseRouterInit(r)
	IssuerRouterInit(r)
}
