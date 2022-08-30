package v1

import (
	"dm/sampleinfo"
	"github.com/jinzhu/copier"
	"rm/common"
	"rm/dto"
	"rm/utils/apollo"
	"rm/utils/thrift"
)

func GetIndexDisplayMaps(platformKey int32) map[string]interface{} {
	return map[string]interface{}{
		"index1DisplayMaps": apollo.BKSampleConfig.GetIndex1DisplayMapByPlatformKey(platformKey),
		"index2DisplayMaps": apollo.BKSampleConfig.GetIndex2DisplayMapByPlatformKey(platformKey),
	}
}

func GetIndex1SelectionMaps(platformKey int32, index1NameSelectionKey int32) []interface{} {
	res := []interface{}{}
	for _, v := range apollo.BKSampleConfig.GetIndex1DisplayMapByPlatformKey(platformKey) {
		if index1NameSelectionKey == 0 || v.Index1NameSelectionKey == index1NameSelectionKey {
			res = append(res, v.IndexMap.Values())
		}
	}
	return res
}

func GetIndex2SelectionMaps(platformKey int32, index2NameSelectionKey int32) []interface{} {
	res := []interface{}{}
	for _, v := range apollo.BKSampleConfig.GetIndex2DisplayMapByPlatformKey(platformKey) {
		if index2NameSelectionKey == 0 || v.Index2NameSelectionKey == index2NameSelectionKey {
			res = append(res, v.IndexMap.Values())
		}
	}
	return res
}

func GetRead1Index(platformKey int32) interface{} {
	if p, ok := apollo.BKCommonConfig.GetPlatform(platformKey); ok && p.PlatformIndexMode == 2 {
		return apollo.BKSampleConfig.Read1IndexMaps
	}
	return []interface{}{}
}

func GetSampleInfoList(token string, searchCriteria *sampleinfo.SearchCriteria, orderBy int32) (*common.ResList, error) {
	data, err := thrift.GetSampleInfoList(token, searchCriteria, orderBy)
	if err != nil {
		return nil, err
	}

	length := len(data.GetSampleInfoList())
	list := make([]*dto.SampleInfoResponse, length, length)
	for i, info := range data.GetSampleInfoList() {
		response := dto.NewSampleInfoResponse()
		if err := copier.Copy(response, info); err != nil {
			common.Log.Warn(err)
		}
		if *info.Index1LengthKey > 0 {
			response.Index1 = true
		}
		if *info.Index2LengthKey > 0 {
			response.Index2 = true
		}
		list[i] = response
	}
	return common.NewResList(list, data.GetSampleTotalCount()), nil
}

func GetSampleInfo(token string, id int32) (*dto.SampleInfoResponse, error) {
	info, err := thrift.GetSampleInfo(token, id)
	if err != nil {
		return nil, err
	}
	response := dto.NewSampleInfoResponse()
	if err := copier.Copy(response, info.SampleInfo); err != nil {
		common.Log.Warn(err)
	}
	return response, nil
}

func CompletionSampleInfo(sample *sampleinfo.SampleInfo) {
	if *sample.LibPrepMethodKey != 0 {
		platform, _ := apollo.BKCommonConfig.GetPlatform(*sample.LibPrepMethodKey)
		sample.LibPrepMethod = &platform.PlatformDisplayName
	}
	if *sample.Index1LengthKey != 0 {
		index, _ := apollo.BKCommonConfig.GetIndex1(*sample.Index1LengthKey)
		sample.Index1Length = &index.Index1Length
	}
	if *sample.Index2LengthKey != 0 {
		index, _ := apollo.BKCommonConfig.GetIndex2(*sample.Index2LengthKey)
		sample.Index2Length = &index.Index2Length
	}
}

func AddSampleInfo(token string, request *dto.AddSampleInfoRequest) (int32, error) {
	sample := sampleinfo.NewSampleInfo()
	if err := copier.Copy(sample, request); err != nil {
		common.Log.Error(err)
		return 0, err
	}
	CompletionSampleInfo(sample)

	data, err := thrift.AddSampleInfo(token, sample)
	if err != nil {
		return 0, err
	}
	return data.GetSampleId(), err
}

func UpdateSampleInfo(token string, id int32, request *dto.SampleInfoRequest) error {
	data, err := thrift.GetSampleInfo(token, id)
	if err != nil {
		return err
	}
	sample := data.GetSampleInfo()
	if err := copier.Copy(sample, request); err != nil {
		common.Log.Error(err)
		return err
	}
	CompletionSampleInfo(sample)

	_, err = thrift.UpdateSampleInfo(token, id, sample)
	return err
}

func RemoveSampleInfo(token string, id int32) error {
	_, err := thrift.RemoveSampleInfo(token, id)
	return err
}
