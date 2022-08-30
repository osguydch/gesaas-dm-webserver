package v1

import (
	"dm/runmanagement"
	"dm/runtask"
	"dm/sampleinfo"
	"fmt"
	"github.com/jinzhu/copier"
	"rm/common"
	"rm/dto"
	"rm/utils/apollo"
	"rm/utils/thrift"
)

func GetRunList(token string, searchCriteria *runtask.SearchCriteria, orderBy int32) (*common.ResList, error) {
	data, err := thrift.GetRunList(token, searchCriteria, orderBy)
	if err != nil {
		return nil, err
	}

	length := len(data.GetRunList())
	list := make([]*dto.RunResponse, length, length)
	for i, run := range data.GetRunList() {
		response := dto.NewRunResponse()
		response.ID = run.ID
		if run.IsSetRuntask() {
			response.RunId = run.GetRuntask().GetRunId()
			response.StartTime = run.GetRuntask().GetStartTime()
			response.LastUpdateTime = run.GetRuntask().GetLastUpdateTime()
			response.OperatorName = run.GetRuntask().GetOperatorName()
			if run.GetRuntask().IsSetRunplan() {
				response.Runname = run.GetRuntask().GetRunplan().GetName()
				response.ThroughputMode = run.GetRuntask().GetRunplan().GetThroughputMode()
				response.ReadlengthMode = run.GetRuntask().GetRunplan().GetReadlengthMode()
				response.AnalysisPipeline = run.GetRuntask().GetRunplan().GetAnalysisPipeline()
			} else {
				common.Log.Warnf("runId: %v, not found run plan", run.ID)
			}
		} else {
			common.Log.Warnf("runId: %v, not found run task", run.ID)
		}
		response.TotalSpaceOccupied = run.GetTotalSpaceOccupied()
		response.StatusDisplayName = common.MetadataConfig.GetRunStatusDisplayName(common.RunStatus(run.Runtask.Status))
		list[i] = response
	}

	return common.NewResList(list, data.GetTotalCount()), nil
}

func GetRunTask(token, runID string) (*dto.RunTaskResponse, error) {
	data, err := thrift.GetRunTask(token, runID)
	if err != nil {
		return nil, err
	}
	if len(data.GetRunTaskList()) == 0 {
		return nil, common.NotExists
	}
	task := data.GetRunTaskList()[0]

	response := dto.NewRunTaskResponse()

	response.DeviceName = task.GetDeviceName()
	response.StartTime = task.GetStartTime()
	response.OperatorName = task.GetOperatorName()
	if task.IsSetRunplan() {
		response.Runname = task.GetRunplan().GetName()
		response.CreatorName = task.GetRunplan().GetCreatorName()
		response.LibPrepMethod = task.GetRunplan().GetLibPrepMethod()
		response.ThroughputMode = task.GetRunplan().GetThroughputMode()
		response.AnalysisPipeline = task.GetRunplan().GetAnalysisPipeline()
		response.ReadlengthMode = task.GetRunplan().GetReadlengthMode()
		response.Reference = apollo.BKRunPlanConfig.GetReferenceDisplayName(task.GetRunplan().GetReferenceKey())
		response.Recipe = task.GetRunplan().GetRecipe()
		response.ScanMatrix = task.GetRunplan().GetScanMatrix()
		response.TrimAdapter = "否"
		if task.GetRunplan().GetIsTrimAdapter() {
			response.TrimAdapter = "是"
		}
		sampleTag := ""
		if task.GetRunplan().GetIndex1LengthKey() > 0 {
			sampleTag += fmt.Sprintf("index1 %vbp", *task.GetRunplan().Index1Length)
		}
		if task.GetRunplan().GetIndex2LengthKey() > 0 {
			sampleTag += fmt.Sprintf("  index2 %vbp", *task.GetRunplan().Index2Length)
		}
		response.SampleTag = sampleTag
		response.LibPrepMethodKey = task.GetRunplan().GetLibPrepMethodKey()
		response.Index1LengthKey = task.GetRunplan().GetIndex1LengthKey()
		response.Index1Length = task.GetRunplan().GetIndex1Length()
		if response.Index1LengthKey > 0 {
			response.Index1 = true
		}
		response.Index2LengthKey = task.GetRunplan().GetIndex2LengthKey()
		response.Index2Length = task.GetRunplan().GetIndex2Length()
		if response.Index2LengthKey > 0 {
			response.Index2 = true
		}
	}
	if task.IsSetChipInfo() {
		response.ChipInfoNo = task.GetChipInfo().GetID()
	}
	if task.IsSetWboxInfo() {
		response.WboxInfoNo = task.GetWboxInfo().GetID()
	}
	if task.IsSetSboxInfo() {
		response.SboxInfoNo = task.GetSboxInfo().GetID()
	}
	if task.IsSetSampleinfo() && task.GetSampleinfo().IsSetID() {
		response.SampleId = task.GetSampleinfo().GetID()
		response.SampleName = task.GetSampleinfo().GetName()
		response.Samples = task.GetSampleinfo().GetSamples()
	} else {
		response.Samples = []*sampleinfo.SingleSampleInfo{}
	}
	return response, nil
}

func getAnalysisIds(infos []*runmanagement.RunDataInfo) []string {
	res := []string{}
	for _, info := range infos {
		res = append(res, info.GetRunResult_().GetAnalysisId())
	}
	return res
}

func GetRunResult(token, runId string) (*dto.RunResultResponse, error) {
	data, err := thrift.GetRun(token, runId)
	if err != nil {
		return nil, err
	}
	if len(data.GetRunList()) == 0 {
		return nil, common.NotExists
	}
	run := data.GetRunList()[0]

	response := dto.NewRunResultResponse()
	if len(run.GetRunDataInfoList()) > 0 {
		runResult := run.GetRunDataInfoList()[0].GetRunResult_()
		if runResult.IsSetAmplificationSummary() {
			response.LoadingRateMean = runResult.GetAmplificationSummary().GetLoadingRateMean()
			response.LoadingImage = runResult.GetAmplificationSummary().GetLoadingImage()
		}
		if runResult.IsSetRunSummary() {
			if runResult.GetRunSummary().IsSetRead1() {
				response.Read1Summary = *runResult.GetRunSummary().GetRead1()
			}
			if runResult.GetRunSummary().IsSetRead2() {
				response.Read2 = true
				response.Read2Summary = *runResult.GetRunSummary().GetRead2()
			}
			response.QcResult = runResult.GetRunSummary().GetQcResult_()
		}
		if runResult.IsSetReferenceSummary() {
			response.RefRatio = runResult.GetReferenceSummary().GetRefRatio()
		}
	}
	response.AnalysisIds = getAnalysisIds(run.GetRunDataInfoList())
	return response, nil
}

func GetAnalysis(token, analysisId string) (*dto.AnalysisResponse, error) {
	data, err := thrift.GetRunResult(token, analysisId)
	if err != nil {
		return nil, err
	}

	response := dto.NewAnalysisResponse()
	response.SaveDir = data.GetRunresult().GetSaveDir()
	response.TotalSpaceOccupied = data.GetRunresult().GetTotalSpaceOccupied()
	if data.GetRunresult().IsSetSampleinfo() {
		response.SampleName = data.GetRunresult().GetSampleinfo().GetName()
		response.LibPrepMethod = data.GetRunresult().GetSampleinfo().GetLibPrepMethod()
		response.Samples = data.GetRunresult().GetSampleinfo().GetSamples()
	} else {
		response.Samples = []*sampleinfo.SingleSampleInfo{}
	}
	response.LibrarySummary = data.GetRunresult().GetLibrarySummary()
	return response, nil
}

func RemoveRun(token, runId string) error {
	_, err := thrift.DeleteRun(token, runId)
	return err
}

func DeleteAnalysis(token, analysisId string) error {
	_, err := thrift.DeleteAnalysis(token, analysisId)
	return err
}

func StartOfflineAnalysis(token string, request *dto.OfflineAnalysisRequest) (string, error) {
	sample := sampleinfo.NewSampleInfo()
	if err := copier.Copy(sample, request.GetSampleInfoRequest()); err != nil {
		common.Log.Error(err)
		return "", err
	}
	CompletionSampleInfo(sample)

	resp, err := thrift.StartOfflineAnalysis(token, *request.RunId, sample, *request.IsTrimAdapter)
	if err != nil {
		return "", err
	}
	return *resp.Analysisid, nil
}
