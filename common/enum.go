package common

type PlanType int32

//测序类型
const (
	NormalPlan PlanType = 1 //普通测序
	CustomPlan PlanType = 2 //高级测序
)

type RunStatus int32

//run状态
const (
	Finish      RunStatus = 0  //完成
	Planned     RunStatus = 1  //已计划
	Starting    RunStatus = 2  //启动中
	Running     RunStatus = 3  //进行中
	Paused      RunStatus = 4  //暂停
	Error       RunStatus = 5  //故障中
	Terminating RunStatus = 6  //外部中止中
	Aborting    RunStatus = 7  //异常中止中
	Terminated  RunStatus = 8  //外部中止
	Aborted     RunStatus = 9  //内部中止
	Deleting    RunStatus = 20 //删除中
	Deleted     RunStatus = 21 //已删除
	DeleteError RunStatus = 22 //删除失败
	Reanalyzing RunStatus = 23 //重分析中
)
