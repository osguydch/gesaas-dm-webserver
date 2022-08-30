package dto

import (
	"dm/runresult"
	"dm/sampleinfo"
	"github.com/jinzhu/copier"
	"rm/common"
)

type RunResponse struct {
	ID                 int32   `json:"id"`
	RunId              string  `json:"runId"`
	Runname            string  `json:"runname"`
	StartTime          int64   `json:"startTime"`
	LastUpdateTime     int64   `json:"lastUpdateTime"`
	ThroughputMode     string  `json:"throughputMode"`
	ReadlengthMode     string  `json:"readlengthMode"`
	AnalysisPipeline   string  `json:"analysisPipeline"`
	OperatorName       string  `json:"operatorName"`
	TotalSpaceOccupied float64 `json:"totalSpaceOccupied"`
	StatusDisplayName  string  `json:"statusDisplayName"`
}

func NewRunResponse() *RunResponse {
	return &RunResponse{}
}

type RunTaskResponse struct {
	Runname          string                         `json:"runName"`
	DeviceName       string                         `json:"deviceName"`
	StartTime        int64                          `json:"startTime"`
	CreatorName      string                         `json:"creatorName"`
	OperatorName     string                         `json:"operatorName"`
	LibPrepMethod    string                         `json:"libPrepMethod"`
	ThroughputMode   string                         `json:"throughputMode"`
	AnalysisPipeline string                         `json:"analysisPipeline"`
	ReadlengthMode   string                         `json:"readlengthMode"`
	Reference        string                         `json:"reference"`
	Recipe           string                         `json:"recipe"`
	ScanMatrix       string                         `json:"scanMatrix"`
	TrimAdapter      string                         `json:"trimAdapter"`
	SampleTag        string                         `json:"sampleTag"`
	ChipInfoNo       int32                          `json:"chipInfoNo"`
	WboxInfoNo       int32                          `json:"wboxInfoNo"`
	SboxInfoNo       int32                          `json:"sboxInfoNo"`
	LibPrepMethodKey int32                          `json:"libPrepMethodKey"`
	Index1           bool                           `json:"index1"`
	Index1LengthKey  int32                          `json:"index1LengthKey"`
	Index1Length     int32                          `json:"index1Length"`
	Index2           bool                           `json:"index2"`
	Index2LengthKey  int32                          `json:"index2LengthKey"`
	Index2Length     int32                          `json:"index2Length"`
	SampleId         int32                          `json:"sampleId"`
	SampleName       string                         `json:"sampleName"`
	Samples          []*sampleinfo.SingleSampleInfo `json:"samples"`
}

func NewRunTaskResponse() *RunTaskResponse {
	return &RunTaskResponse{}
}

type RunResultResponse struct {
	LoadingRateMean float64               `json:"loadingRateMean"`
	Read1Summary    runresult.ReadSummary `json:"read1Summary"`
	Read2           bool                  `json:"read2"`
	Read2Summary    runresult.ReadSummary `json:"read2Summary"`
	QcResult        int32                 `json:"qcResult"`
	RefRatio        float64               `json:"refRatio"`
	LoadingImage    string                `json:"loadingImage"`
	AnalysisIds     []string              `json:"analysisIds"`
}

func NewRunResultResponse() *RunResultResponse {
	return &RunResultResponse{}
}

type AnalysisResponse struct {
	SaveDir            string                         `json:"saveDir"`
	TotalSpaceOccupied float64                        `json:"totalSpaceOccupied"`
	SampleName         string                         `json:"sampleName"`
	LibPrepMethod      string                         `json:"libPrepMethod"`
	Samples            []*sampleinfo.SingleSampleInfo `json:"samples"`
	LibrarySummary     []*runresult.LibraryInfo       `json:"librarySummary"`
}

func NewAnalysisResponse() *AnalysisResponse {
	return &AnalysisResponse{}
}

type OfflineAnalysisRequest struct {
	RunId            *string                        `json:"runId" binding:"required"`
	IsTrimAdapter    *bool                          `json:"isTrimAdapter" binding:"required"`
	SampleName       *string                        `json:"sampleName"`
	LibPrepMethodKey *int32                         `json:"libPrepMethodKey"`
	Index1           *bool                          `json:"index1"`
	Index1LengthKey  *int32                         `json:"index1LengthKey"`
	Index2           *bool                          `json:"index2"`
	Index2LengthKey  *int32                         `json:"index2LengthKey"`
	Samples          []*sampleinfo.SingleSampleInfo `json:"samples"`
}

func NewOfflineAnalysisRequest() *OfflineAnalysisRequest {
	return &OfflineAnalysisRequest{}
}

func (p *OfflineAnalysisRequest) GetSampleInfoRequest() *SampleInfoRequest {
	sampleInfoRequest := NewSampleInfoRequest()
	if err := copier.Copy(sampleInfoRequest, p); err != nil {
		common.Log.Error(err)
		return nil
	}
	sampleInfoRequest.Name = p.SampleName
	return sampleInfoRequest
}

type BathRemoveRunRequest struct {
	RunIDs []string `json:"runIDs" binding:"required"`
}

func NewBathRemoveRunRequest() *BathRemoveRunRequest {
	return &BathRemoveRunRequest{}
}
