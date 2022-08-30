package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

type ResDesc string

var (
	// 全局
	Success           ResDesc = "Success"
	InvalidParamError ResDesc = "InvalidParamError"
	NotExistsError    ResDesc = "NotExistsError"
	NotSupportError   ResDesc = "NotSupportError"
	UnknownError      ResDesc = "UnknownError"

	// dm 错误
	DmFailedError                        ResDesc = "DmFailedError"
	DmInvalidTokenError                  ResDesc = "DmInvalidTokenError"
	DmNoAuthError                        ResDesc = "DmNoAuthError"
	DmInvalidCreatorIdError              ResDesc = "DmInvalidCreatorIdError"
	DmInvalidStatusError                 ResDesc = "DmInvalidStatusError"
	DmRunIsNotFinishedError              ResDesc = "DmRunIsNotFinishedError"
	DmRunIsFinishedError                 ResDesc = "DmRunIsFinishedError"
	DmInvalidSampleIdError               ResDesc = "DmInvalidSampleIdError"
	DmSampleNameAlreadyExistsError       ResDesc = "DmSampleNameAlreadyExistsError"
	DmEmptySampleNameError               ResDesc = "DmEmptySampleNameError"
	DmInvalidRunplanIdError              ResDesc = "DmInvalidRunPlanIdError"
	DmRunplanIdAlreadyExistsError        ResDesc = "DmRunPlanIdAlreadyExistsError"
	DmInvalidRuntaskIdError              ResDesc = "DmInvalidRunTaskIdError"
	DmInvalidAnalysisidError             ResDesc = "DmInvalidAnalysisIdError"
	DmAnalysisidAlreadyExistsError       ResDesc = "DmAnalysisIdAlreadyExistsError"
	DmInvalidRunidError                  ResDesc = "DmInvalidRunIdError"
	DmEmptyAnalysisidError               ResDesc = "DmEmptyAnalysisIdError"
	DmEmptyRunidError                    ResDesc = "DmEmptyRunIdError"
	DmInvalidDeviceNameError             ResDesc = "DmInvalidDeviceNameError"
	DmDeviceNameAlreadyExistsError       ResDesc = "DmDeviceNameAlreadyExistsError"
	DmEmptyDeviceNameError               ResDesc = "DmEmptyDeviceNameError"
	DmInvalidDataTypeError               ResDesc = "DmInvalidDataTypeError"
	DmDataTypeAlreadyExistsError         ResDesc = "DmDataTypeAlreadyExistsError"
	DmDataGroupAlreadyExistsError        ResDesc = "DmDataGroupAlreadyExistsError"
	DmInvalidDataGroupError              ResDesc = "DmInvalidDataGroupError"
	DmDataGroupMappingAlreadyExistsError ResDesc = "DmDataGroupMappingAlreadyExistsError"
	DmEmptyGroupNameError                ResDesc = "DmEmptyGroupNameError"
	DmInvalidProjectNameError            ResDesc = "DmInvalidProjectNameError"
	DmProjectNameAlreadyExistsError      ResDesc = "DmProjectNameAlreadyExistsError"
	DmInvalidRunDataError                ResDesc = "DmInvalidRunDataError"
	DmRunDataAlreadyExistsError          ResDesc = "DmRunDataAlreadyExistsError"
	DmInvalidMacAddressError             ResDesc = "DmInvalidMacAddressError"
	DmMacAddressAlreadyExistsError       ResDesc = "DmMacAddressAlreadyExistsError"
	DmEmptyMacAddressError               ResDesc = "DmEmptyMacAddressError"

	// 业务 错误
	InvalidIdError                ResDesc = "InvalidIdError"
	FieldTypeError                ResDesc = "FieldTypeError"
	MissingSampleError            ResDesc = "MissingSampleError"
	TooMuchSampleError            ResDesc = "TooMuchSampleError"
	SampleNameInvalidError        ResDesc = "SampleNameInvalidError"
	SampleNameReservedError       ResDesc = "SampleNameReservedError"
	SampleNameExistsError         ResDesc = "SampleNameExistsError"
	Read1IndexInvalidError        ResDesc = "Read1IndexInvalidError"
	IndexExistsError              ResDesc = "IndexExistsError"
	Index1SeqLengthError          ResDesc = "Index1SeqLengthError"
	Index1SeqInvalidError         ResDesc = "Index1SeqInvalidError"
	Index1MismatchPredefinedError ResDesc = "Index1MismatchPredefinedError"
	Index2SeqLengthError          ResDesc = "Index2SeqLengthError"
	Index2SeqInvalidError         ResDesc = "Index2SeqInvalidError"
	Index2MismatchPredefinedError ResDesc = "Index2MismatchPredefinedError"
	Index1InvalidError            ResDesc = "Index1InvalidError"
	Index2InvalidError            ResDesc = "Index2InvalidError"
	Read1IndexKeyExistsError      ResDesc = "Read1IndexKeyExistsError"
	Read1IndexSeqInvalidError     ResDesc = "Read1IndexSeqInvalidError"
	Read1IndexSeqExistsError      ResDesc = "Read1IndexSeqExistsError"
)

var (
	OK = &ResStatus{Code: "0x00000000", Description: Success, Solution: ""}
	//InvalidToken = &ResStatus{Code: "0x28010101", Description: InvalidTokenError, Solution: "请先登录"}

	InvalidParam = &ResStatus{Code: "0x28110001", Description: InvalidParamError, Solution: ""}
	NotExists    = &ResStatus{Code: "0x28110002", Description: NotExistsError, Solution: ""}
	NotSupport   = &ResStatus{Code: "0x28110003", Description: NotSupportError, Solution: ""}

	Unknown = &ResStatus{Code: "0x99999999", Description: UnknownError, Solution: "服务异常，请联系管理员"}
)

func helper(code string, msg ResDesc) *ResStatus {
	return &ResStatus{
		Code:        code,
		Description: msg,
		Solution:    "",
	}
}

type ResStatus struct {
	Code        string        `json:"code"`
	Description ResDesc       `json:"description"`
	Params      []interface{} `json:"-"`
	Solution    string        `json:"solution"`
}

func (p *ResStatus) Error() string {
	return string(p.Description)
}

func (p *ResStatus) Desc(msg ResDesc, params ...interface{}) *ResStatus {
	return &ResStatus{p.Code, msg, params, p.Solution}
}

func (p *ResStatus) SetParam(params ...interface{}) *ResStatus {
	return &ResStatus{p.Code, p.Description, params, p.Solution}
}

func (p *ResStatus) Translator(trans ut.Translator) *ResStatus {
	params := []string{}
	if len(p.Params) > 0 {
		for _, param := range p.Params {
			params = append(params, fmt.Sprintf("%v", param))
		}
	}
	t, err := trans.T(string(p.Description), params...)
	if err != nil {
		Log.Warnf("translator %s err: %v", p.Description, err)
		t = string(p.Description)
	}
	return &ResStatus{p.Code, ResDesc(t), p.Params, p.Solution}
}

func (p *ResStatus) H(data ...interface{}) gin.H {
	if len(data) > 0 {
		return gin.H{"status": helper(p.Code, p.Description), "data": data[0]}
	}
	return gin.H{"status": helper(p.Code, p.Description), "data": gin.H{}}
}
