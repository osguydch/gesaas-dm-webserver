package v1

import (
	"dm/runplan"
	"dm/sampleinfo"
	"github.com/jinzhu/copier"
	"rm/common"
	"rm/dto"
	"rm/utils/apollo"
	"rm/utils/thrift"
)

func GetThroughputModeDefinitions(platformKey int32) interface{} {
	return apollo.BKRunPlanConfig.GetThroughputDefinitionsByPlatformKey(platformKey)
}

func GetReadlengthModeDefinitions(platformKey, throughputModeKey int32) interface{} {
	definitions, _ := apollo.BKRunPlanConfig.GetReadlengthModeDefinitionsByPlatformKey(platformKey, throughputModeKey)
	return definitions
}

func GetReferenceMaps(platformKey int32) interface{} {
	return apollo.BKRunPlanConfig.GetReferences(platformKey)
}

func GetAnalysisMaps() interface{} {
	return common.MetadataConfig.Analysis
}

func GetBkTypes(platformKey int32) interface{} {
	return apollo.BKConsumableConfig.GetTypes(platformKey)
}

func GetBkVariants(platformKey int32, bkType string) interface{} {
	variants, _ := apollo.BKConsumableConfig.GetVariants(platformKey, bkType)
	return variants
}

func GetBkVersions(platformKey int32, bkType, bkVariant string) interface{} {
	versions, _ := apollo.BKConsumableConfig.GetVersions(platformKey, bkType, bkVariant)
	return versions
}

func GetRunPlanList(token string, searchCriteria *runplan.SearchCriteria, orderBy int32) (*common.ResList, error) {
	data, err := thrift.GetRunPlanList(token, searchCriteria, orderBy)
	if err != nil {
		return nil, err
	}

	length := len(data.GetRunPlanList())
	list := make([]*dto.RunPlanResponse, length, length)
	for i, plan := range data.GetRunPlanList() {
		response := dto.NewRunPlanResponse()
		if err := copier.Copy(response, plan); err != nil {
			common.Log.Warn(err)
		}
		if *plan.Index1LengthKey > 0 {
			response.Index1 = true
		}
		if *plan.Index2LengthKey > 0 {
			response.Index2 = true
		}
		if *plan.ReferenceKey > 0 {
			response.Ref = true
			reference, ok := apollo.BKRunPlanConfig.GetReference(*plan.ReferenceKey)
			if !ok {
				common.Log.Warn(err)
			} else {
				response.Reference = reference.ReferenceDisplayName
			}
		}
		if *plan.SampleId > 0 {
			sample, err := GetSampleInfo(token, *plan.SampleId)
			if err != nil {
				common.Log.Warn(err)
				response.Samples = []*dto.SingleSampleInfo{}
			} else {
				response.SampleName = sample.Name
				response.Samples = sample.Samples
			}
		} else {
			response.Samples = []*dto.SingleSampleInfo{}
		}
		list[i] = response
	}
	return common.NewResList(list, data.GetRunPlanTotalCount()), nil
}

func GetRunPlan(token string, id int32) (*dto.RunPlanResponse, error) {
	data, err := thrift.GetRunPlan(token, id)
	if err != nil {
		return nil, err
	}

	response := dto.NewRunPlanResponse()
	if err := copier.Copy(response, data.GetRunPlan()); err != nil {
		common.Log.Warn(err)
	}
	if response.Index1LengthKey > 0 {
		response.Index1 = true
	}
	if response.Index2LengthKey > 0 {
		response.Index2 = true
	}
	if response.ReferenceKey > 0 {
		response.Ref = true
		reference, ok := apollo.BKRunPlanConfig.GetReference(response.ReferenceKey)
		if !ok {
			common.Log.Warn(err)
		} else {
			response.Reference = reference.ReferenceDisplayName
		}
	}
	if response.SampleId > 0 {
		sample, err := GetSampleInfo(token, response.SampleId)
		if err != nil {
			common.Log.Warn(err)
			response.Samples = []*dto.SingleSampleInfo{}
		} else {
			response.SampleName = sample.Name
			response.Samples = sample.Samples
		}
	} else {
		response.Samples = []*dto.SingleSampleInfo{}
	}
	return response, nil
}

func CompletionRunPlan(plan *runplan.RunPlan) {
	platform, _ := apollo.BKCommonConfig.GetPlatform(*plan.LibPrepMethodKey)
	plan.LibPrepMethod = &platform.PlatformDisplayName
	throughput, _ := apollo.BKRunPlanConfig.GetThroughputDefinition(*plan.ThroughputModeKey)
	plan.ThroughputMode = &throughput.ThroughputModeName
	readlength, _ := apollo.BKRunPlanConfig.GetReadlengthModeDefinition(*plan.ReadlengthModeKey)
	plan.ReadlengthMode = &readlength.ReadlengthModeDisplayName

	if *plan.Index1LengthKey != 0 {
		index, _ := apollo.BKCommonConfig.GetIndex1(*plan.Index1LengthKey)
		plan.Index1Length = &index.Index1Length
	}
	if *plan.Index2LengthKey != 0 {
		index, _ := apollo.BKCommonConfig.GetIndex2(*plan.Index2LengthKey)
		plan.Index2Length = &index.Index2Length
	}
	if *plan.ReferenceKey != 0 {
		detail, _ := apollo.BKRunPlanConfig.GetReferenceDetail(*plan.ReferenceKey)
		plan.Reference = sampleinfo.NewSingleSampleInfo()
		if err := copier.Copy(plan.Reference, detail); err != nil {
			common.Log.Error(err)
		}
	}
	plan.IsDisplay = true
}

func AddRunPlan(token string, request *dto.AddRunPlanRequest) (int32, error) {
	plan := runplan.NewRunPlan()
	if err := copier.Copy(plan, request); err != nil {
		common.Log.Error(err)
		return 0, err
	}
	CompletionRunPlan(plan)

	if *plan.SampleId > 0 {
		if err := UpdateSampleInfo(token, *plan.SampleId, request.GetSampleInfoRequest()); err != nil {
			return 0, err
		}
	}
	data, err := thrift.AddRunPlan(token, plan)
	if err != nil {
		return 0, err
	}
	return data.GetRunplanId(), nil
}

func UpdateRunPlan(token string, id int32, request *dto.RunPlanRequest) error {
	data, err := thrift.GetRunPlan(token, id)
	if err != nil {
		return err
	}
	plan := data.GetRunPlan()
	if err := copier.Copy(plan, request); err != nil {
		common.Log.Error(err)
		return err
	}
	CompletionRunPlan(plan)

	if err := UpdateSampleInfo(token, plan.GetSampleId(), request.GetSampleInfoRequest()); err != nil {
		common.Log.Error(err)
		return err
	}

	_, err = thrift.UpdateRunPlan(token, id, plan)
	return err
}

func RemoveRunPlan(token string, id int32) error {
	_, err := thrift.RemoveRunPlan(token, id)
	return err
}
