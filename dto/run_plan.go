package dto

import (
	"github.com/jinzhu/copier"
	"rm/common"
)

type RunPlanRequest struct {
	Name              *string             `json:"name" binding:"required,runPlanNameValidator"`
	Type              *int32              `json:"type" binding:"required,runPlanTypeValidator"`
	LibPrepMethodKey  *int32              `json:"libPrepMethodKey" binding:"required,platformKeyValidator"`
	ThroughputModeKey *int32              `json:"throughputModeKey" binding:"required,runPlanThroughputModeValidator"`
	ReadlengthModeKey *int32              `json:"readlengthModeKey" binding:"required,runPlanReadlengthModeValidator"`
	Index1            *bool               `json:"index1" binding:"required"`
	Index1LengthKey   *int32              `json:"index1LengthKey" binding:"required,index1LengthKeyValidator"`
	Index2            *bool               `json:"index2" binding:"required"`
	Index2LengthKey   *int32              `json:"index2LengthKey" binding:"required,index2LengthKeyValidator"`
	AnalysisPipeline  *string             `json:"analysisPipeline" binding:"required,runPlanAnalysisValidator"`
	Ref               *bool               `json:"ref" binding:"required"`
	ReferenceKey      *int32              `json:"referenceKey" binding:"required,runPlanReferenceValidator"`
	SampleId          *int32              `json:"sampleId" binding:"required"`
	SampleName        *string             `json:"sampleName" binding:"required"`
	Samples           []*SingleSampleInfo `json:"samples" binding:"required,dive"`
	IsTrimAdapter     *bool               `json:"isTrimAdapter" binding:"required"`
	BkType            *string             `json:"bkType" binding:"required"`
	BkVariant         *string             `json:"bkVariant" binding:"required"`
	BkVersion         *string             `json:"bkVersion" binding:"required"`
}

func NewRunPlanRequest() *RunPlanRequest {
	return &RunPlanRequest{}
}

func (p *RunPlanRequest) GetSampleInfoRequest() *SampleInfoRequest {
	sampleInfoRequest := NewSampleInfoRequest()
	if err := copier.Copy(sampleInfoRequest, p); err != nil {
		common.Log.Error(err)
		return nil
	}
	sampleInfoRequest.Name = p.SampleName
	return sampleInfoRequest
}

type AddRunPlanRequest struct {
	RunPlanRequest
	CreatorId *int32 `json:"creatorId" binding:"required"`
}

func NewAddRunPlanRequest() *AddRunPlanRequest {
	return &AddRunPlanRequest{}
}

type RunPlanResponse struct {
	ID                  int32               `json:"id"`
	Name                string              `json:"name"`
	Type                int32               `json:"type"`
	LibPrepMethodKey    int32               `json:"libPrepMethodKey"`
	LibPrepMethod       string              `json:"libPrepMethod"`
	ThroughputModeKey   int32               `json:"throughputModeKey"`
	ThroughputMode      string              `json:"throughputMode"`
	ReadlengthModeKey   int32               `json:"readlengthModeKey"`
	ReadlengthMode      string              `json:"readlengthMode"`
	Index1              bool                `json:"index1"`
	Index1LengthKey     int32               `json:"index1LengthKey"`
	Index1Length        int32               `json:"index1Length"`
	Index2              bool                `json:"index2"`
	Index2LengthKey     int32               `json:"index2LengthKey"`
	Index2Length        int32               `json:"index2Length"`
	AnalysisPipelineKey int32               `json:"analysisPipelineKey"`
	AnalysisPipeline    string              `json:"analysisPipeline"`
	Ref                 bool                `json:"ref"`
	ReferenceKey        int32               `json:"referenceKey"`
	Reference           string              `json:"reference"`
	SampleId            int32               `json:"sampleId"`
	SampleName          string              `json:"sampleName"`
	Samples             []*SingleSampleInfo `json:"samples"`
	CreatorName         string              `json:"creatorName"`
	CreateTime          int32               `json:"createTime"`
	LastUpdateTime      int32               `json:"lastUpdateTime"`
	IsTrimAdapter       bool                `json:"isTrimAdapter"`
	BkType              string              `json:"bkType"`
	BkVariant           string              `json:"bkVariant"`
	BkVersion           string              `json:"bkVersion"`
	IsDisplay           bool                `json:"isDisplay"`
}

func NewRunPlanResponse() *RunPlanResponse {
	return &RunPlanResponse{}
}
