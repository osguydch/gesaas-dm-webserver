package thrift

import (
	"dm/runplan"
	"rm/common"
)

func GetRunPlanList(token string, searchCriteria *runplan.SearchCriteria, orderBy int32) (*runplan.ReturnData, error) {
	tClient, err := newThriftClient(runplan.RunPlanManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runplan.NewRunPlanManagementClient(tClient)
	resp, err := client.GetRunPlanList(dfCtx, token, searchCriteria, 10, orderBy)

	common.Log.Debugf("searchCriteria: %s, resp %s, err: %v", common.P(searchCriteria), common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

func GetRunPlan(token string, id int32) (*runplan.ReturnData, error) {
	tClient, err := newThriftClient(runplan.RunPlanManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runplan.NewRunPlanManagementClient(tClient)
	resp, err := client.GetRunPlan(dfCtx, token, id)

	common.Log.Debugf("id %v, resp %s, err: %v", id, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

func AddRunPlan(token string, plan *runplan.RunPlan) (*runplan.ReturnData, error) {
	tClient, err := newThriftClient(runplan.RunPlanManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runplan.NewRunPlanManagementClient(tClient)
	resp, err := client.AddRunPlan(dfCtx, token, plan)

	defer common.Log.Debugf("plan: %s, resp: %s, err: %v", common.P(plan), common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}

func UpdateRunPlan(token string, id int32, plan *runplan.RunPlan) (*runplan.ReturnData, error) {
	tClient, err := newThriftClient(runplan.RunPlanManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runplan.NewRunPlanManagementClient(tClient)
	resp, err := client.UpdateRunPlan(dfCtx, token, id, plan)

	defer common.Log.Debugf("id: %v, plan: %s, resp: %s, err: %v", id, common.P(plan), common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}

func RemoveRunPlan(token string, id int32) (*runplan.ReturnData, error) {
	tClient, err := newThriftClient(runplan.RunPlanManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := runplan.NewRunPlanManagementClient(tClient)
	resp, err := client.RemoveRunPlan(dfCtx, token, id)

	defer common.Log.Debugf("id: %v, resp: %s, err: %v", id, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}
