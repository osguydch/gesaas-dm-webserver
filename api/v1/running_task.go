package v1

import (
	"dm/runningtask"
	"github.com/gin-gonic/gin"
	"rm/common"
	"rm/dto"
	service "rm/service/v1"
	"rm/utils"
)

func GetTaskTypes(c *gin.Context) {
	resp := service.GetTaskTypes()

	utils.HttpSuccess(c, resp)
}

func GetRunningTaskList(c *gin.Context) {
	runName, deviceName, typeName := c.Query("runName"), c.Query("deviceName"), c.Query("type")
	pageNum, pageSize := common.Pagination(c)
	searchCriteria := &runningtask.SearchCriteria{
		PageNum:  &pageNum,
		PageSize: &pageSize,
	}
	if len(runName) > 0 {
		searchCriteria.Runname = &runName
	}
	if len(deviceName) > 0 {
		searchCriteria.DeviceName = &deviceName
	}
	if len(typeName) > 0 {
		taskType, err := runningtask.TaskTypeFromString(typeName)
		if err != nil {
			utils.HttpErr(c, common.InvalidParam.SetParam("任务类型"))
			return
		}
		searchCriteria.Type = &taskType
	}

	resp, err := service.GetRunningTaskList(token(c), searchCriteria, common.OrderBy(c))
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func StopAnalysis(c *gin.Context) {
	request := dto.NewStopAnalysisRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.HttpErr(c, err)
		return
	}
	if err := service.StopAnalysis(token(c), request.AnalysisId); err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c)
}
