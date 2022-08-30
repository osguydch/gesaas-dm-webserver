package dto

type SingleSampleInfo struct {
	SampleID     *string `json:"sampleID" binding:"required"`
	Description  *string `json:"description" binding:"required"`
	Index1ID     *string `json:"index1ID" binding:"required"`
	Index1       *string `json:"index1" binding:"required"`
	Index2ID     *string `json:"index2ID" binding:"required"`
	Index2       *string `json:"index2" binding:"required"`
	Read1IndexID *string `json:"read1IndexID" binding:"required"`
	Read1Index   *string `json:"read1Index" binding:"required"`
	RefGenome    *string `json:"refGenome" binding:"required"`
}

type SampleInfoRequest struct {
	Name             *string             `json:"name" binding:"required,sampleNameValidator"`
	LibPrepMethodKey *int32              `json:"libPrepMethodKey" binding:"required,platformKeyValidator"`
	Index1           *bool               `json:"index1" binding:"required"`
	Index1LengthKey  *int32              `json:"index1LengthKey" binding:"required,index1LengthKeyValidator"`
	Index2           *bool               `json:"index2" binding:"required"`
	Index2LengthKey  *int32              `json:"index2LengthKey" binding:"required,index2LengthKeyValidator"`
	Samples          []*SingleSampleInfo `json:"samples" binding:"required,dive"`
}

func NewSampleInfoRequest() *SampleInfoRequest {
	return &SampleInfoRequest{}
}

type AddSampleInfoRequest struct {
	SampleInfoRequest
	CreatorId *int32 `json:"creatorId" binding:"required"`
}

func NewAddSampleInfoRequest() *AddSampleInfoRequest {
	return &AddSampleInfoRequest{}
}

type SampleInfoResponse struct {
	ID               int32               `json:"id"`
	Name             string              `json:"name"`
	LibPrepMethodKey int32               `json:"libPrepMethodKey"`
	LibPrepMethod    string              `json:"libPrepMethod"`
	Index1           bool                `json:"index1"`
	Index1LengthKey  int32               `json:"index1LengthKey"`
	Index1Length     int32               `json:"index1Length"`
	Index2           bool                `json:"index2"`
	Index2LengthKey  int32               `json:"index2LengthKey"`
	Index2Length     int32               `json:"index2Length"`
	Samples          []*SingleSampleInfo `json:"samples"`
	CreatorName      string              `json:"creatorName"`
	CreateTime       int32               `json:"createTime"`
	LastUpdateTime   int32               `json:"lastUpdateTime"`
}

func NewSampleInfoResponse() *SampleInfoResponse {
	return &SampleInfoResponse{}
}
