package v1

import (
	"dm/sampleinfo"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"rm/common"
	"rm/dto"
	service "rm/service/v1"
	"rm/utils"
	"rm/utils/apollo"
	"rm/utils/validation"
)

// GetIndexDisplayMaps
// 建库方法与Index序列展示映射关系以及Index序列展示的大类名称
func GetIndexDisplayMaps(c *gin.Context) {
	platformKey := c.Query("platformKey")

	resp := service.GetIndexDisplayMaps(cast.ToInt32(platformKey))

	utils.HttpSuccess(c, resp)
}

// GetIndex1SelectionMaps
// Index1序列展示定义和具体的Index1名称和序列之间的映射关系
func GetIndex1SelectionMaps(c *gin.Context) {
	platformKey := c.Query("platformKey")
	index1NameSelectionKey := c.Query("index1NameSelectionKey")

	res := service.GetIndex1SelectionMaps(cast.ToInt32(platformKey), cast.ToInt32(index1NameSelectionKey))

	utils.HttpSuccess(c, res)
}

// GetIndex2SelectionMaps
// Index2序列展示定义和具体的Index2名称和序列之间的映射关系
func GetIndex2SelectionMaps(c *gin.Context) {
	platformKey := c.Query("platformKey")
	index2NameSelectionKey := c.Query("index2NameSelectionKey")

	res := service.GetIndex2SelectionMaps(cast.ToInt32(platformKey), cast.ToInt32(index2NameSelectionKey))

	utils.HttpSuccess(c, res)
}

// GetRead1Index
// 样本信息中Read1Index的名称和序列
func GetRead1Index(c *gin.Context) {
	platformKey := c.Query("platformKey")

	res := service.GetRead1Index(cast.ToInt32(platformKey))

	utils.HttpSuccess(c, res)
}

// GetGeneGroupDefinition
// 样本信息基因组的配置
func GetGeneGroupDefinition(c *gin.Context) {
	utils.HttpSuccess(c, apollo.BKSampleConfig.GeneGroups)
}

func GetSampleInfoList(c *gin.Context) {
	sampleName, creatorName := c.Query("sampleName"), c.Query("creatorName")
	libPrepMethod, index1Length, index2Length := c.Query("libPrepMethod"), c.Query("index1Length"), c.Query("index2Length")
	pageNum, pageSize := common.Pagination(c)
	searchCriteria := &sampleinfo.SearchCriteria{
		PageNum:  &pageNum,
		PageSize: &pageSize,
	}
	if len(sampleName) > 0 {
		searchCriteria.SampleName = &sampleName
	}
	if len(creatorName) > 0 {
		searchCriteria.CreatorName = &creatorName
	}
	if len(libPrepMethod) > 0 {
		searchCriteria.LibPrepMethod = &libPrepMethod
	}
	if len(index1Length) > 0 {
		index1Len, err := cast.ToInt32E(index1Length)
		if err != nil {
			utils.HttpErr(c, common.InvalidParam.SetParam("index1Length"))
			return
		}
		searchCriteria.Index1Length = &index1Len
	}
	if len(index2Length) > 0 {
		index2Len, err := cast.ToInt32E(index2Length)
		if err != nil {
			utils.HttpErr(c, common.InvalidParam.SetParam("index2Length"))
			return
		}
		searchCriteria.Index2Length = &index2Len
	}

	resp, err := service.GetSampleInfoList(token(c), searchCriteria, common.OrderBy(c))
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func AddSampleInfo(c *gin.Context) {
	request := dto.NewAddSampleInfoRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.HttpErr(c, err)
		return
	}
	if err := validation.ValidateSampleInfoRequest(&request.SampleInfoRequest); err != nil {
		utils.HttpErr(c, err)
		return
	}

	resp, err := service.AddSampleInfo(token(c), request)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func GetSampleInfo(c *gin.Context) {
	id := cast.ToUint32(c.Param("id"))
	if id == 0 {
		utils.HttpErr(c, common.InvalidParam.Desc(common.InvalidIdError))
		return
	}

	resp, err := service.GetSampleInfo(token(c), int32(id))
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c, resp)
}

func UpdateSampleInfo(c *gin.Context) {
	id := cast.ToUint32(c.Param("id"))
	if id == 0 {
		utils.HttpErr(c, common.InvalidParam.Desc(common.InvalidIdError))
		return
	}

	request := dto.NewSampleInfoRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.HttpErr(c, err)
		return
	}
	if err := validation.ValidateSampleInfoRequest(request); err != nil {
		utils.HttpErr(c, err)
		return
	}

	err := service.UpdateSampleInfo(token(c), int32(id), request)
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c)
}

func RemoveSampleInfo(c *gin.Context) {
	id := cast.ToUint32(c.Param("id"))
	if id == 0 {
		utils.HttpErr(c, common.InvalidParam.Desc(common.InvalidIdError))
		return
	}

	err := service.RemoveSampleInfo(token(c), int32(id))
	if err != nil {
		utils.HttpErr(c, err)
		return
	}
	utils.HttpSuccess(c)
}
