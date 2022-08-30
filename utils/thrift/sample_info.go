package thrift

import (
	"dm/sampleinfo"
	"rm/common"
)

func GetSampleInfoList(token string, searchCriteria *sampleinfo.SearchCriteria, orderBy int32) (*sampleinfo.ReturnData, error) {
	tClient, err := newThriftClient(sampleinfo.SampleManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := sampleinfo.NewSampleManagementClient(tClient)
	resp, err := client.GetSampleInfoList(dfCtx, token, searchCriteria, 10, orderBy)

	common.Log.Debugf("searchCriteria: %s, resp %s, err: %v", common.P(searchCriteria), common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

func GetSampleInfo(token string, id int32) (*sampleinfo.ReturnData, error) {
	tClient, err := newThriftClient(sampleinfo.SampleManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := sampleinfo.NewSampleManagementClient(tClient)
	resp, err := client.GetSampleInfo(dfCtx, token, id)

	common.Log.Debugf("id %v, resp %s, err: %v", id, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

func AddSampleInfo(token string, sample *sampleinfo.SampleInfo) (*sampleinfo.ReturnData, error) {
	tClient, err := newThriftClient(sampleinfo.SampleManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := sampleinfo.NewSampleManagementClient(tClient)
	resp, err := client.AddSampleInfo(dfCtx, token, sample)

	defer common.Log.Debugf("sample: %s, resp: %s, err: %v", common.P(sample), common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}

func UpdateSampleInfo(token string, id int32, sample *sampleinfo.SampleInfo) (*sampleinfo.ReturnData, error) {
	tClient, err := newThriftClient(sampleinfo.SampleManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := sampleinfo.NewSampleManagementClient(tClient)
	resp, err := client.UpdateSampleInfo(dfCtx, token, id, sample)

	defer common.Log.Debugf("id: %v, sample: %s, resp: %s, err: %v", id, common.P(sample), common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}

func RemoveSampleInfo(token string, id int32) (*sampleinfo.ReturnData, error) {
	tClient, err := newThriftClient(sampleinfo.SampleManagementName)
	if err != nil {
		return nil, err
	}
	defer tClient.Close()

	client := sampleinfo.NewSampleManagementClient(tClient)
	resp, err := client.RemoveSampleInfo(dfCtx, token, id)

	defer common.Log.Debugf("id: %v, resp: %s, err: %v", id, common.P(resp), err)

	if err := errorHandling(resp, err); err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}
