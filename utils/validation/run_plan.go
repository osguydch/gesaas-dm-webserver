package validation

import (
	"rm/common"
	"rm/dto"
	"rm/utils/apollo"
)

func ValidateRunPlanRequest(request *dto.RunPlanRequest) error {
	definitions, ok := apollo.BKRunPlanConfig.GetReadlengthModeDefinitionsByPlatformKey(*request.LibPrepMethodKey, *request.ThroughputModeKey)
	if !ok {
		return common.InvalidParam.SetParam("通量模式")
	}
	validReadlengthModeKey := false
	for _, definition := range definitions {
		if definition.ReadlengthModeKey == *request.ReadlengthModeKey {
			validReadlengthModeKey = true
		}
	}
	if !validReadlengthModeKey {
		return common.InvalidParam.SetParam("读长模式")
	}

	platform, _ := apollo.BKCommonConfig.GetPlatform(*request.LibPrepMethodKey)
	if platform.PlatformIndexMode == 1 {
		if *request.Index1 == false && *request.Index2 == true {
			return common.InvalidParam.SetParam("Index设置不正确")
		}
		if *request.Index1 == true && *request.SampleId == 0 {
			return common.InvalidParam.SetParam("样本表设置不正确")
		}
		if *request.Index1 == false && *request.SampleId != 0 {
			return common.InvalidParam.SetParam("样本表设置不正确")
		}
	}
	if platform.PlatformIndexMode == 2 {
		if *request.Index1 != false || *request.Index2 != false {
			return common.InvalidParam.SetParam("Index设置不正确")
		}
		if *request.SampleId == 0 {
			return common.InvalidParam.SetParam("样本表设置不正确")
		}
	}

	if *request.Ref == true {
		if *request.ReferenceKey == 0 {
			return common.InvalidParam.SetParam("内参")
		}
		reference, _ := apollo.BKRunPlanConfig.GetReference(*request.ReferenceKey)
		if reference.PlatformKey != platform.PlatformKey {
			return common.InvalidParam.SetParam("内参")
		}
	}
	if *request.Ref == false && *request.ReferenceKey != 0 {
		return common.InvalidParam.SetParam("内参")
	}

	if *request.SampleId > 0 {
		if err := ValidateSampleInfoRequest(request.GetSampleInfoRequest()); err != nil {
			return err
		}
	}

	if *request.Type == int32(common.CustomPlan) {
		//高级测序模式，检查BK类型
		_, ok = apollo.BKConsumableConfig.GetVariants(platform.PlatformKey, *request.BkType)
		if !ok {
			return common.InvalidParam.SetParam("耗材Type")
		}
		_, ok = apollo.BKConsumableConfig.GetVersions(platform.PlatformKey, *request.BkType, *request.BkVariant)
		if !ok {
			return common.InvalidParam.SetParam("耗材Variant")
		}
		ok = apollo.BKConsumableConfig.ContainsVersion(platform.PlatformKey, *request.BkType, *request.BkVariant, *request.BkVersion)
		if !ok {
			return common.InvalidParam.SetParam("耗材Version")
		}
	}
	return nil
}
