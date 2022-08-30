package dto

import "dm/runningtask"

type TaskInfo struct {
	RemainingTime       float64                   `json:"remainingTime"`
	LoadingRate         runningtask.IndicatorItem `json:"loadingRate"`
	Throughput          runningtask.IndicatorItem `json:"throughput"`
	Aq30                runningtask.IndicatorItem `json:"aq30"`
	SealRate            runningtask.IndicatorItem `json:"sealRate"`
	EstimatedFinishTime float64                   `json:"estimatedFinishTime"`
}

type RunningTaskResponse struct {
	AnalysisId      string               `json:"analysisId"`
	RunId           string               `json:"runId"`
	DeviceName      string               `json:"deviceName"`
	Runname         string               `json:"runname"`
	Type            runningtask.TaskType `json:"type"`
	TypeDisplayName string               `json:"typeDisplayName"`
	CreatorName     string               `json:"creatorName"`
	Status          int32                `json:"status"`
	TaskInfo        TaskInfo             `json:"taskInfo"`
	CreatedAt       float64              `json:"createdAt"`
	UpdatedAt       float64              `json:"updatedAt"`
}

func NewRunningTaskResponse() *RunningTaskResponse {
	return &RunningTaskResponse{}
}

type StopAnalysisRequest struct {
	AnalysisId string `json:"analysisID" binding:"required"`
}

func NewStopAnalysisRequest() *StopAnalysisRequest {
	return &StopAnalysisRequest{}
}
