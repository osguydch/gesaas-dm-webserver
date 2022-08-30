package thrift

import (
	"dm/runresult"
	"rm/common"
)

func GetRunResult(token, analysisId string) (*runresult.ReturnData, error) {
	tClient, err := newThriftClient(runresult.RunResultManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runresult.NewRunResultManagementClient(tClient)
	resp, err := client.GetRunResult_(dfCtx, token, analysisId)

	common.Log.Debugf("id: %s, resp %s, err: %v", analysisId, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
