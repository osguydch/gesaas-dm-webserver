package validation

import (
	"github.com/emirpasic/gods/sets/hashset"
	"regexp"
	"rm/common"
	"rm/dto"
	"rm/utils/apollo"
)

var (
	sampleIdRe    *regexp.Regexp
	sampleIdWords = hashset.New("Unknown", "Undetermined", "Control", "KeyMismatched", "ShortRead")
	indexRe       *regexp.Regexp
	read1IndexRe  *regexp.Regexp
)

func init() {
	sampleIdRe, _ = regexp.Compile(`^[0-9A-Za-z][0-9A-Za-z-_]{0,98}$`)
	indexRe, _ = regexp.Compile(`^[ATCG]+$`)
	read1IndexRe, _ = regexp.Compile(`^[ATCG]{1,20}$`)
}

func ValidateSampleInfoRequest(request *dto.SampleInfoRequest) error {
	platform, _ := apollo.BKCommonConfig.GetPlatform(*request.LibPrepMethodKey)
	if platform.PlatformIndexMode == 1 {
		//mode为1时 index1必勾选
		if *request.Index1 == false {
			return common.InvalidParam.SetParam("index1")
		}
		//mode为1时 index1LengthKey必须选择
		if *request.Index1LengthKey == 0 {
			return common.InvalidParam.SetParam("index1LengthKey")
		}
		//mode为1、index2勾选时 index2LengthKey必须选择
		if *request.Index2 == true && *request.Index2LengthKey == 0 {
			return common.InvalidParam.SetParam("index2LengthKey")
		}
		//mode为1、index2不勾选时 index2LengthKey必须不选
		if *request.Index2 == false && *request.Index2LengthKey != 0 {
			return common.InvalidParam.SetParam("index2LengthKey")
		}
	}
	if platform.PlatformIndexMode == 2 {
		//mode为2时 index1禁止勾选
		if *request.Index1 == true {
			return common.InvalidParam.SetParam("index1")
		}
		//mode为2时 index2禁止勾选
		if *request.Index2 == true {
			return common.InvalidParam.SetParam("index2")
		}
		//mode为2时 index1LengthKey必须不选
		if *request.Index1LengthKey != 0 {
			return common.InvalidParam.SetParam("index1LengthKey")
		}
		//mode为2时 index2LengthKey必须不选
		if *request.Index2LengthKey != 0 {
			return common.InvalidParam.SetParam("index2LengthKey")
		}
	}
	//至少包含一个样本
	if len(request.Samples) == 0 {
		return common.InvalidParam.Desc(common.MissingSampleError)
	}
	//最多99个样本
	if len(request.Samples) > 99 {
		return common.InvalidParam.Desc(common.TooMuchSampleError)
	}
	sampleIDs := hashset.New()
	indexs := hashset.New()
	read1IndexIDs, read1Indexs := hashset.New(), hashset.New()
	for i, sample := range request.Samples {
		//sampleID格式
		if !sampleIdRe.MatchString(*sample.SampleID) {
			return common.InvalidParam.Desc(common.SampleNameInvalidError, i+1, sample.SampleID)
		}
		//sampleID不能包含保留字
		if ok := sampleIdWords.Contains(sample.SampleID); ok {
			return common.InvalidParam.Desc(common.SampleNameReservedError, i+1, sample.SampleID)
		}
		//sampleID不能重复
		if ok := sampleIDs.Contains(sample.SampleID); ok {
			return common.InvalidParam.Desc(common.SampleNameExistsError, i+1, sample.SampleID)
		}
		sampleIDs.Add(sample.SampleID)

		if platform.PlatformIndexMode == 1 {
			//mode为1时 read1Index不能有值
			if len(*sample.Read1IndexID) != 0 || len(*sample.Read1Index) != 0 {
				return common.InvalidParam.Desc(common.Read1IndexInvalidError, i+1)
			}
			index1, _ := apollo.BKCommonConfig.GetIndex1(*request.Index1LengthKey)
			//mode为1时 index1长度不能超过Index1Length设置
			if len(*sample.Index1) > int(index1.Index1Length) {
				return common.InvalidParam.Desc(common.Index1SeqLengthError, i+1, sample.Index1, index1.Index1Length)
			}
			//mode为1时 index1格式
			if !indexRe.MatchString(*sample.Index1) {
				return common.InvalidParam.Desc(common.Index1SeqInvalidError, i+1, sample.Index1)
			}
			//mode为1时 index1如果已预定义 必须跟index1ID匹配
			if index1ID, ok := apollo.BKSampleConfig.I1IndexMapContains(*request.LibPrepMethodKey, *sample.Index1); ok {
				if *sample.Index1ID != index1ID {
					return common.InvalidParam.Desc(common.Index1MismatchPredefinedError, i+1, sample.Index1)
				}
			}

			if *request.Index2 == true {
				index2, _ := apollo.BKCommonConfig.GetIndex2(*request.Index2LengthKey)
				//mode为1、index2被勾选时 index2长度不能超过Index2Length设置
				if len(*sample.Index2) > int(index2.Index2Length) {
					return common.InvalidParam.Desc(common.Index2SeqLengthError, i+1, sample.Index2, index2.Index2Length)
				}
				//mode为1、index2被勾选时 index2格式
				if !indexRe.MatchString(*sample.Index2) {
					return common.InvalidParam.Desc(common.Index2SeqInvalidError, i+1, sample.Index2)
				}
				//mode为1、index2被勾选时 index2如果已预定义 必须跟index2ID匹配
				if index2ID, ok := apollo.BKSampleConfig.I2IndexMapContains(*request.LibPrepMethodKey, *sample.Index2); ok {
					if *sample.Index2ID != index2ID {
						return common.InvalidParam.Desc(common.Index2MismatchPredefinedError, i+1, sample.Index2)
					}
				}
			}
			//mode为1时 index1 + index2不能重复
			if ok := indexs.Contains(*sample.Index1 + "-" + *sample.Index2); ok {
				return common.InvalidParam.Desc(common.IndexExistsError, i+1)
			}
			indexs.Add(*sample.Index1 + "-" + *sample.Index2)
		}
		if platform.PlatformIndexMode == 2 {
			//mode为2时 index1ID、index1不能有值
			if len(*sample.Index1ID) != 0 || len(*sample.Index1) != 0 {
				return common.InvalidParam.Desc(common.Index1InvalidError, i+1)
			}
			//mode为2时 index2ID、index2不能有值
			if len(*sample.Index2ID) != 0 || len(*sample.Index2) != 0 {
				return common.InvalidParam.Desc(common.Index2InvalidError, i+1)
			}

			if len(*sample.Read1IndexID) > 0 {
				//mode为2时 read1IndexID可以不选但不能重复
				if ok := read1IndexIDs.Contains(sample.Read1IndexID); ok {
					return common.InvalidParam.Desc(common.Read1IndexKeyExistsError, i+1, sample.Read1IndexID)
				}
				read1IndexIDs.Add(sample.Read1IndexID)
			}
			//mode为2时 read1Index格式
			if !read1IndexRe.MatchString(*sample.Read1Index) {
				return common.InvalidParam.Desc(common.Read1IndexSeqInvalidError, i+1, sample.Read1Index)
			}
			//mode为2时 read1Index不能重复
			if ok := read1Indexs.Contains(sample.Read1Index); ok {
				return common.InvalidParam.Desc(common.Read1IndexSeqExistsError, i+1, sample.Read1Index)
			}
			read1Indexs.Add(sample.Read1Index)
		}
	}

	return nil
}
