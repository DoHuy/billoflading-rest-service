package ginh

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"custom-webhook-store-logs/packages/logger/zapl"
)

func LogInfo(ctx *gin.Context, msg string, fields ...zap.Field) {
	zapl.Info(RequestContextFromGinCtx(ctx), msg, fields...)
}

func LogError(ctx *gin.Context, msg string, err error, fields ...zap.Field) {
	zapl.Error(RequestContextFromGinCtx(ctx), msg, err, fields...)
}
