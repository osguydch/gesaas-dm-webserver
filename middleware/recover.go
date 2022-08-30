package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"rm/common"
	"rm/utils"
	"runtime/debug"
)

func GinRecovery(logger *zap.SugaredLogger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.Errorf("[Recovery from panic] %v stack %s", recovered, string(debug.Stack()))
		utils.HttpErr(c, common.Unknown)
		c.AbortWithStatus(http.StatusOK)
	})
}
