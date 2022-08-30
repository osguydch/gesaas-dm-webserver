package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GinLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()

		cost := time.Since(start)
		logger.Infof("[%s]%s, ip[%s], resp[%d] %s errors[%s]",
			c.Request.Method,
			path,
			c.ClientIP(),
			c.Writer.Status(),
			cost.String(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
