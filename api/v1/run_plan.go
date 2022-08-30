package v1

import (
	"dm/runplan"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"rm/common"
	"rm/dto"
	service "rm/service/v1"
	"rm/utils"
	"rm/utils/validation"
)

// GetThroughputModeDefinitions
// 测序计划通量模式配置
func GetThroughputModeDefinitions(c *gin.Context) {
	platformKey := c.Query("platformKey")

	resp := service.GetThroughputModeDefinitions(cast.ToInt32(platformKey))

	utils.HttpSuccess(c, resp)
}

// GetReadlengthModeDefinitions
// 测序计划读长模式配置
func GetReadlengthModeDefinitions(c *gin.Context) {
	platformKey := c.Query("platformKey")
	throughputModeKey := c.Query("throughputModeKey")

	resp := service.GetReadlengthModeDefinitions(cast.ToInt32(platformKey), cast.ToInt32(throughputModeKey))

	utils.HttpSuccess(c, resp)
}

// GetReferenceMaps
// 测序计划内参选项配置
func GetReferenceMaps(c *gin.Context) {
	platformKey := c.Query("platformKey")

	resp := service.GetReferenceMaps(cast.ToInt32(platformKey))

	utils.HttpSuccess(c, resp)
}

// GetAnalysisMaps
// 测序计划分析流程选项配置
func GetAnalysisMaps(c *gin.Context) {

	resp := service.GetAnalysisMaps()

	utils.HttpSuccess(c, resp)
}

// GetBkTypes
// 测序计划耗材版本Type选项配置
func GetBkTypes(c *gin.Context) {
	platformKey := c.Query("platformKey")

	resp := service.GetBkTypes(cast.ToInt32(platformKey))

	utils.HttpSuccess(c, resp)
}

// GetBkVariants
// 测序计划耗材版本Variant选项配置
func GetBkVariants(c *gin.Context) {
	platformKey := c.Query("platformKey")
	bkType := c.Query("bkType")

	resp := service.GetBkVariants(cast.ToInt32(platformKey), bkType)

	utils.HttpSuccess(c, resp)
}

// GetBkVersions
// 测序计划耗材版本Version选项配置
func GetBkVersions(c *gin.Context) {
	platformKey := c.Query("platformKey")
	bkType := c.Query("bkType")
	bkVariant := c.Query("bkVariant")

	resp := service.GetBkVersions(cast.ToInt32(platformKey), bkType, bkVariant)

	utils.HttpSuccess(c, resp)
}

func GetRunPlanList(c *gin.Context) {
	planName, creatorName := c.Query("planName"), c.Query("creatorName")
	pageNum, pageSize := common.Pagination(c)
	searchCriteria := &runplan.SearchCriteria{
		PageNum:  &pageNum,
		PageSize: &pageSize,
	}
	if len(planName) > 0 {
		searchCriteria.Name = &planName
	}
	if len(creatorName) > 0 {
		searchCriteria.CreatorName = &creatorName
	}

	resp, err := service.GetRunPlanList(token(c), searchCriteria, common.OrderBy(c))
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func AddRunPlan(c *gin.Context) {
	request := dto.NewAddRunPlanRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.HttpErr(c, err)
		return
	}

	if err := validation.ValidateRunPlanRequest(&request.RunPlanRequest); err != nil {
		utils.HttpErr(c, err)
		return
	}

	resp, err := service.AddRunPlan(token(c), request)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func UpdateRunPlan(c *gin.Context) {
	id := cast.ToUint32(c.Param("id"))
	if id == 0 {
		utils.HttpErr(c, common.InvalidParam.Desc(common.InvalidIdError))
		return
	}

	request := dto.NewRunPlanRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.HttpErr(c, err)
		return
	}

	if err := validation.ValidateRunPlanRequest(request); err != nil {
		utils.HttpErr(c, err)
		return
	}

	err := service.UpdateRunPlan(token(c), int32(id), request)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c)
}

func RemoveRunPlan(c *gin.Context) {
	id := cast.ToUint32(c.Param("id"))
	if id == 0 {
		utils.HttpErr(c, common.InvalidParam.Desc(common.InvalidIdError))
		return
	}

	err := service.RemoveRunPlan(token(c), int32(id))
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c)
}
