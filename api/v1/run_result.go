package v1

import (
	"dm/runtask"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"rm/common"
	"rm/dto"
	service "rm/service/v1"
	"rm/utils"
	"rm/utils/validation"
)

// GetRunStatusMaps
// 测序结果状态选项
func GetRunStatusMaps(c *gin.Context) {
	utils.HttpSuccess(c, common.MetadataConfig.RunStatus.Values())
}

func GetRunList(c *gin.Context) {
	runName, creatorName, status := c.Query("runName"), c.Query("creatorName"), c.Query("status")
	pageNum, pageSize := common.Pagination(c)
	searchCriteria := &runtask.SearchCriteria{
		PageNum:  &pageNum,
		PageSize: &pageSize,
	}
	if len(runName) > 0 {
		searchCriteria.Runname = &runName
	}
	if len(creatorName) > 0 {
		searchCriteria.Creator = &creatorName
	}
	if len(status) > 0 {
		s, err := cast.ToInt32E(status)
		if err != nil {
			utils.HttpErr(c, common.InvalidParam.SetParam("status"))
			return
		}
		searchCriteria.Status = &s
	}

	resp, err := service.GetRunList(token(c), searchCriteria, common.OrderBy(c))
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func BatchRemoveRun(c *gin.Context) {
	request := dto.NewBathRemoveRunRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.HttpErr(c, err)
		return
	}

	for _, runID := range request.RunIDs {
		err := service.RemoveRun(token(c), runID)
		if err != nil {
			utils.HttpErr(c, err)
			return
		}
	}
	utils.HttpSuccess(c)
}

func GetRunTask(c *gin.Context) {
	runID := c.Param("id")

	resp, err := service.GetRunTask(token(c), runID)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func GetRunResult(c *gin.Context) {
	runId := c.Param("id")

	resp, err := service.GetRunResult(token(c), runId)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func GetAnalysis(c *gin.Context) {
	analysisId := c.Param("id")

	resp, err := service.GetAnalysis(token(c), analysisId)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func RemoveAnalysis(c *gin.Context) {
	analysisId := c.Param("id")

	err := service.DeleteAnalysis(token(c), analysisId)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c)
}

func StartOfflineAnalysis(c *gin.Context) {
	request := dto.NewOfflineAnalysisRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.HttpErr(c, err)
		return
	}

	task, err := service.GetRunTask(token(c), *request.RunId)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}

	request.LibPrepMethodKey = &task.LibPrepMethodKey
	request.Index1 = &task.Index1
	request.Index1LengthKey = &task.Index1LengthKey
	request.Index2 = &task.Index2
	request.Index2LengthKey = &task.Index2LengthKey
	if task.SampleId == 0 {
		//当实时分析不选样本表时不能重分析
		utils.HttpErr(c, common.NotSupport)
		return
	}

	if err := validation.ValidateSampleInfoRequest(request.GetSampleInfoRequest()); err != nil {
		utils.HttpErr(c, err)
		return
	}
	analysisId, err := service.StartOfflineAnalysis(token(c), request)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, analysisId)
}
