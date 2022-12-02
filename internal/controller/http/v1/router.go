// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	usecase "nextclan/validator-register/mobile-validator-register-service/internal/usecase"
	"nextclan/validator-register/mobile-validator-register-service/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean TempLast API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, cmv usecase.CreateMobileValidatorDevice, umv usecase.UpdateMobileValidatorDeviceStatus, pa usecase.PusherAuthentication, l logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	//Library for API docs
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	//how well is the http server running
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	//
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newDeviceRoutes(h, cmv, umv, pa, l)
	}
}