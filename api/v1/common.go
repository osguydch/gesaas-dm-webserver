package v1

import (
	"github.com/gin-gonic/gin"
	"rm/common"
	"rm/utils"
	"rm/utils/apollo"
)

func token(c *gin.Context) string {
	return "token-value: " + c.GetHeader("token")
}

func GetGlobal(c *gin.Context) {
	utils.HttpSuccess(c, struct {
		GetTaskInterval int64 `json:"getTaskInterval"`
	}{
		common.MetadataConfig.GetTaskInterval,
	})
}

// GetCommon
// 建库平台/方法配置
// index1
// index2
func GetCommon(c *gin.Context) {
	utils.HttpSuccess(c, apollo.BKCommonConfig)
}
