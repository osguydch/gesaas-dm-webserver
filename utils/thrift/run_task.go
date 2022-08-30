package thrift

import (
	"dm/runtask"
	"rm/common"
)

func GetRunTask(token, runId string) (*runtask.ReturnData, error) {
	tClient, err := newThriftClient(runtask.RunTaskManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runtask.NewRunTaskManagementClient(tClient)
	resp, err := client.GetRunTask(dfCtx, token, runId)

	common.Log.Debugf("id: %s, resp %s, err: %v", runId, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
