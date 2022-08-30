package v1

import (
	"dm/runningtask"
	"github.com/jinzhu/copier"
	"rm/common"
	"rm/dto"
	"rm/utils/thrift"
)

func GetTaskTypes() []interface{} {
	return common.MetadataConfig.TaskType.Values()
}

func GetRunningTaskList(token string, searchCriteria *runningtask.SearchCriteria, orderBy int32) (*common.ResList, error) {
	data, err := thrift.GetRunningTaskList(token, searchCriteria, orderBy)
	if err != nil {
		return nil, err
	}

	length := len(data.GetTaskList())
	list := make([]*dto.RunningTaskResponse, length, length)
	for i, task := range data.GetTaskList() {
		response := dto.NewRunningTaskResponse()
		if err := copier.Copy(response, task); err != nil {
			common.Log.Warn(err)
		}
		response.TypeDisplayName = common.MetadataConfig.GetTaskTypeDisplayName(task.Type)
		list[i] = response
	}

	return common.NewResList(list, data.GetTotalCount()), nil
}

func StopAnalysis(token, analysisId string) error {
	_, err := thrift.StopAnalysis(token, analysisId)
	if err != nil {
		return err
	}
	return nil
}
