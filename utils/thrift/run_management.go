package thrift

import (
	"dm/runmanagement"
	"dm/runtask"
	"dm/sampleinfo"
	"rm/common"
)

func DeleteRun(token, runId string) (*runmanagement.Response, error) {
	tClient, err := newThriftClient(runmanagement.RunManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runmanagement.NewRunManagementClient(tClient)
	resp, err := client.DeleteRun(dfCtx, token, runId)

	defer common.Log.Debugf("id: %v, resp: %s, err: %v", runId, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteAnalysis(token, analysisId string) (*runmanagement.Response, error) {
	tClient, err := newThriftClient(runmanagement.RunManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runmanagement.NewRunManagementClient(tClient)
	resp, err := client.DeleteAnalysis(dfCtx, token, analysisId)

	defer common.Log.Debugf("id: %v, resp: %s, err: %v", analysisId, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp, nil
}

func GetRunList(token string, searchCriteria *runtask.SearchCriteria, orderBy int32) (*runmanagement.Response, error) {
	tClient, err := newThriftClient(runmanagement.RunManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runmanagement.NewRunManagementClient(tClient)
	resp, err := client.GetRunList(dfCtx, token, searchCriteria, 10, orderBy)

	common.Log.Debugf("searchCriteria: %s, resp %s, err: %v", common.P(searchCriteria), common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp, nil
}

func GetRun(token, runId string) (*runmanagement.Response, error) {
	tClient, err := newThriftClient(runmanagement.RunManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runmanagement.NewRunManagementClient(tClient)
	resp, err := client.GetRun(dfCtx, token, runId)

	common.Log.Debugf("id: %s, resp %s, err: %v", runId, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp, nil
}

func StartOfflineAnalysis(token, runId string, sampleInfo *sampleinfo.SampleInfo, isTrimAdapter bool) (*runmanagement.Response, error) {
	tClient, err := newThriftClient(runmanagement.RunManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runmanagement.NewRunManagementClient(tClient)
	analysisParam := map[string]*runmanagement.GenericField{"trimAdapter": {B: &isTrimAdapter}}
	resp, err := client.StartOfflineAnalysis(dfCtx, token, runId, sampleInfo, analysisParam)

	defer common.Log.Debugf("runId: %v, sampleInfo %v, analysisParam %v, resp: %s, err: %v", runId, sampleInfo, analysisParam, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp, nil
}

func StopAnalysis(token, analysisId string) (*runmanagement.Response, error) {
	tClient, err := newThriftClient(runmanagement.RunManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runmanagement.NewRunManagementClient(tClient)
	resp, err := client.StopAnalysis(dfCtx, token, analysisId)

	defer common.Log.Debugf("id: %v, resp: %s, err: %v", analysisId, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp, nil
}
