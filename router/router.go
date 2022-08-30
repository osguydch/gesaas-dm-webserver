package router

import (
	"github.com/gin-gonic/gin"
	apiv1 "rm/api/v1"
	"rm/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	middleware.InitMiddleware(r)

	// 注册路由
	v1 := r.Group("/api/v0.1/rm")
	{
		config := v1.Group("/config")
		{
			//获取配置
			config.GET("/global", apiv1.GetGlobal)
			config.GET("/common", apiv1.GetCommon)
			//样本表配置信息
			config.GET("/index/display", apiv1.GetIndexDisplayMaps)
			config.GET("/index1/selection", apiv1.GetIndex1SelectionMaps)
			config.GET("/index2/selection", apiv1.GetIndex2SelectionMaps)
			config.GET("/read1/selection", apiv1.GetRead1Index)
			config.GET("/gene", apiv1.GetGeneGroupDefinition)
			//测序计划配置信息
			config.GET("/throughput", apiv1.GetThroughputModeDefinitions)
			config.GET("/readlength", apiv1.GetReadlengthModeDefinitions)
			config.GET("/reference", apiv1.GetReferenceMaps)
			config.GET("/analysis", apiv1.GetAnalysisMaps)
			config.GET("/bk/type", apiv1.GetBkTypes)
			config.GET("/bk/variant", apiv1.GetBkVariants)
			config.GET("/bk/version", apiv1.GetBkVersions)
			//测序任务配置信息
			config.GET("/task/type", apiv1.GetTaskTypes)
			//测序结果配置信息
			config.GET("/run/status", apiv1.GetRunStatusMaps)
		}

		//测序计划列表
		v1.GET("/runplans", apiv1.GetRunPlanList)
		//新增测序计划
		v1.POST("/runplans", apiv1.AddRunPlan)
		//修改测序计划
		v1.POST("/runplans/:id", apiv1.UpdateRunPlan)
		//删除测序计划
		v1.DELETE("/runplans/:id", apiv1.RemoveRunPlan)

		//样本表列表
		v1.GET("/sampleinfos", apiv1.GetSampleInfoList)
		//新增样本表
		v1.POST("/sampleinfos", apiv1.AddSampleInfo)
		//获取样本表信息
		v1.GET("/sampleinfos/:id", apiv1.GetSampleInfo)
		//修改样本表
		v1.POST("/sampleinfos/:id", apiv1.UpdateSampleInfo)
		//删除样本表
		v1.DELETE("/sampleinfos/:id", apiv1.RemoveSampleInfo)

		//测序任务列表
		v1.GET("/runningtasks", apiv1.GetRunningTaskList)
		//测序任务停止
		v1.POST("/runningtasks/stop", apiv1.StopAnalysis)

		//测序结果列表
		v1.GET("/runs", apiv1.GetRunList)
		//批量删除测序结果
		v1.POST("/runs/batch/delete", apiv1.BatchRemoveRun)
		//获取任务详情
		v1.GET("/runs/:id/task", apiv1.GetRunTask)
		//获取任务结果详情
		v1.GET("/runs/:id/result", apiv1.GetRunResult)
		//获取任务结果列表
		v1.GET("/runs/analysis/:id", apiv1.GetAnalysis)
		//测序结果删除分析结果
		v1.DELETE("/runs/analysis/:id", apiv1.RemoveAnalysis)
		//重分析
		v1.POST("/runs/analysis/offline", apiv1.StartOfflineAnalysis)
	}
	return r
}
