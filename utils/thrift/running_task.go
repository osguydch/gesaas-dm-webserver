package thrift

import (
	"dm/runningtask"
	"rm/common"
)

func GetRunningTaskList(token string, searchCriteria *runningtask.SearchCriteria, orderBy int32) (*runningtask.ReturnData, error) {
	tClient, err := newThriftClient(runningtask.RunningTaskManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runningtask.NewRunningTaskManagementClient(tClient)
	resp, err := client.GetTaskList(dfCtx, token, searchCriteria, orderBy)

	common.Log.Debugf("searchCriteria: %s, resp %s, err: %v", common.P(searchCriteria), common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
