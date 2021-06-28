package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	// for doc
	_ "vtp-apis/docs"
	"vtp-apis/external/delivery/rest"
	commonConf "vtp-apis/packages/config"
	"vtp-apis/packages/constants"
	"vtp-apis/packages/ginh"
	"vtp-apis/packages/logger/ginl"
)

type GinDependencies struct {
	RunMode          commonConf.RunModeConfig
	HealthCheck      *ginh.BaseSerializer
	WebHookSerialize *rest.CustomWebHooksExecutionSerializer
}

func NewGin(deps *GinDependencies) *gin.Engine {
	h := gin.New()

	h.Use(ginl.GinWithZap(zap.L(), time.RFC3339, true))

	customHooks := h.Group("/web-hooks")
	{
		health := customHooks.Group("/health")
		{
			health.GET("", deps.HealthCheck.Health)
		}
		if deps.RunMode.RunMode == constants.RunModeNonProduction {
			fmt.Println("deps.RunMode.RunMode: ", deps.RunMode.RunMode)
			customHooks.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		customHooks.POST("/logs/incoming", deps.WebHookSerialize.IncomingActivityLogsFromAuth0)

	}
	return h
}
