package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"rm/common"
	"rm/router"
	"rm/utils"
	"rm/utils/apollo"
	"rm/utils/thrift"
	"strings"
	"time"
)

var (
	configFolder string
	port         int
	mode         string
	StartCmd     = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "rmserver server -c config/config.toml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.Flags().StringVarP(&configFolder, "config", "c", "config/", "Start server with provided configuration folder")
	StartCmd.Flags().IntVarP(&port, "port", "p", 8000, "Tcp port server listening on")
	StartCmd.Flags().StringVarP(&mode, "mode", "m", "debug", "server mode ; eg:debug,test,release")
	viper.BindPFlag("port", StartCmd.Flags().Lookup("port"))
	viper.BindPFlag("mode", StartCmd.Flags().Lookup("mode"))
}

func setup() {
	if strings.HasSuffix(configFolder, "/") {
		configFolder = configFolder[:len(configFolder)-1]
		if strings.HasSuffix(configFolder, "/") {
			panic("invalid config folder")
		}
	}
	configToml := configFolder + "/config.toml"
	//1. 读取配置
	common.SetUpConfig(configToml)
	//2. 设置日志
	common.SetupLogger()
	//3. 初始化数据库链接
	//database.Setup()
	//4. 初始化thrift
	thrift.SetupThrift()
	//5. 初始化翻译器
	utils.SetupTrans(configFolder)
	//5. 初始化验证器
	utils.SetupValidator()
	//5. 初始化Apollo
	apollo.SetupApollo()

	common.Log.Info(`starting api server`)
	common.Log.Infof("pid: %d\n", os.Getpid())
}

func run() error {
	gin.SetMode(common.ApplicationConfig.Mode)
	r := router.InitRouter()

	srv := &http.Server{
		Addr:    common.ApplicationConfig.Addr(),
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			common.Log.Fatal("listen: ", err)
		}
	}()

	fmt.Println("Server run at:")
	fmt.Printf("-  Network: http://%s:%d/ \n", common.ApplicationConfig.Host, common.ApplicationConfig.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \n", common.Now())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \n", common.Now())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		common.Log.Fatal("Server Shutdown:", err)
	}
	common.Log.Info("Server exiting")
	common.ShutdownLogger()
	return nil
}
