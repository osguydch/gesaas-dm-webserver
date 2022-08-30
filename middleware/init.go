package middleware

import (
	"github.com/gin-gonic/gin"
	"rm/common"
)

func InitMiddleware(r *gin.Engine) {
	// NoCache is a middleware function that appends headers
	r.Use(NoCache)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)
	// Use Zap Logger
	r.Use(GinLogger(common.Log))
	// Global Recover
	r.Use(GinRecovery(common.Log))
}
