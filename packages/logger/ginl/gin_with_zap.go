package ginl

import (
	"time"
	"vtp-apis/packages/logger/zapl"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinWithZap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
func GinWithZap(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		contextFields := zapl.BuildFieldsWithContext(c, []zap.Field{zap.String("log_type", "gin")})
		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e, contextFields...)
			}
		} else {
			contextFields = append(contextFields,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
			)
			logger.Info(path, contextFields...)
		}
	}
}
