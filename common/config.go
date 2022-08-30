package common

import (
	"dm/runningtask"
	"fmt"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/spf13/viper"
)

//配置文件

var (
	ApplicationConfig    = new(application)
	LoggerConfig         = new(logger)
	DataManagementConfig = new(dataManagement)
	ApolloConfig         = new(apollo)
	MetadataConfig       = new(metadata)
)

// 应用配置
type application struct {
	Host string
	Port int
	Mode string
}

func (p *application) Addr() string {
	return fmt.Sprintf("%s:%d", p.Host, p.Port)
}

// 日志配置
type logger struct {
	LogDir   string // 日志文件夹路径
	LogLevel string // 日志打印等级
	MaxSize  int    //在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxAge   int    //保留旧文件的最大天数
	Compress bool   //是否压缩/归档旧文件
	Stdout   bool
}

// DM配置
type dataManagement struct {
	Ip        string
	Port      int
	Transport string
	Protocol  string
	Usessl    bool
}

// apollo配置
type apollo struct {
	AppID        string
	Cluster      string
	MetaAddr     string
	BKCommon     string
	BKSample     string
	BKRunPlan    string
	BKConsumable string
}

type analysis struct {
	AnalysisPipelineKey         string `json:"analysisPipelineKey"`
	AnalysisPipelineDisplayName string `json:"analysisPipelineDisplayName"`
	AnalysisPipeline            string `json:"analysisPipeline"`
}

type runStatusDisplayName struct {
	RunStatusKey         RunStatus `json:"runStatusKey"`
	RunStatusDisplayName string    `json:"runStatusDisplayName"`
}

type taskTypeDisplayName struct {
	TaskTypeKey         runningtask.TaskType `json:"taskTypeKey"`
	TaskTypeDisplayName string               `json:"taskTypeDisplayName"`
}

// metadata配置
type metadata struct {
	GetTaskInterval int64
	Analysis        []*analysis
	TaskType        *linkedhashmap.Map
	RunStatus       *linkedhashmap.Map

	aMap map[string]*analysis
}

func (p *metadata) setup() {
	p.aMap = map[string]*analysis{}
	for _, analysis := range p.Analysis {
		p.aMap[analysis.AnalysisPipelineKey] = analysis
	}
}

func (p *metadata) GetAnalysis(key string) (*analysis, bool) {
	v, ok := p.aMap[key]
	return v, ok
}

func (p *metadata) GetTaskTypeDisplayName(taskType runningtask.TaskType) string {
	if v, ok := p.TaskType.Get(taskType); ok {
		return v.(taskTypeDisplayName).TaskTypeDisplayName
	}
	return ""
}

func (p *metadata) GetRunStatusDisplayName(runStatus RunStatus) string {
	if v, ok := p.RunStatus.Get(runStatus); ok {
		return v.(runStatusDisplayName).RunStatusDisplayName
	}
	return ""
}

func InitApplication() *application {
	return &application{
		Host: viper.GetString("host"),
		Port: viper.GetInt("port"),
		Mode: viper.GetString("mode"),
	}
}

func InitLog(cfg *viper.Viper) *logger {
	var logger logger
	err := cfg.Unmarshal(&logger)
	if err != nil {
		panic("InitLog err")
	}
	return &logger
}

func InitDataManagement(cfg *viper.Viper) *dataManagement {
	var dataManagement dataManagement
	err := cfg.Unmarshal(&dataManagement)
	if err != nil {
		panic("InitDataManagement err")
	}
	return &dataManagement
}

func InitApollo(cfg *viper.Viper) *apollo {
	var apollo apollo
	err := cfg.Unmarshal(&apollo)
	if err != nil {
		panic("InitApollo err")
	}
	return &apollo
}

func InitMetadata(cfg *viper.Viper) *metadata {
	var metadata metadata
	err := cfg.Unmarshal(&metadata)
	if err != nil {
		panic("InitMetadata err")
	}

	//
	metadata.TaskType = linkedhashmap.New()
	metadata.TaskType.Put(runningtask.TaskType_Run, taskTypeDisplayName{runningtask.TaskType_Run, "测序中"})
	metadata.TaskType.Put(runningtask.TaskType_OfflineAnalysis, taskTypeDisplayName{runningtask.TaskType_OfflineAnalysis, "重分析中"})

	//Run状态定义 http://10.0.32.42:4999/web/#/3?page_id=290
	metadata.RunStatus = linkedhashmap.New()
	metadata.RunStatus.Put(Finish, runStatusDisplayName{Finish, "完成"})
	metadata.RunStatus.Put(Terminated, runStatusDisplayName{Terminated, "外部中止"})
	metadata.RunStatus.Put(Aborted, runStatusDisplayName{Aborted, "内部中止"})
	metadata.RunStatus.Put(Deleting, runStatusDisplayName{Deleting, "删除中"})
	metadata.RunStatus.Put(Deleted, runStatusDisplayName{Deleted, "已删除"})
	metadata.RunStatus.Put(DeleteError, runStatusDisplayName{DeleteError, "删除失败"})
	metadata.RunStatus.Put(Reanalyzing, runStatusDisplayName{Reanalyzing, "重分析中"})

	metadata.setup()
	return &metadata
}

// 载入配置文件
func SetUpConfig(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	ApplicationConfig = InitApplication()

	cfgLog := viper.Sub("logger")
	if cfgLog == nil {
		panic("No found logger in the configuration")
	}
	LoggerConfig = InitLog(cfgLog)

	cfgDataManagement := viper.Sub("servers.datamanagement")
	if cfgDataManagement == nil {
		panic("No found servers.datamanagement in the configuration")
	}
	DataManagementConfig = InitDataManagement(cfgDataManagement)

	cfgApolloConfig := viper.Sub("apollo")
	if cfgApolloConfig == nil {
		panic("No found apollo in the configuration")
	}
	ApolloConfig = InitApollo(cfgApolloConfig)

	metadataApolloConfig := viper.Sub("metadata")
	if metadataApolloConfig == nil {
		panic("No found metadata in the configuration")
	}
	MetadataConfig = InitMetadata(metadataApolloConfig)
}
