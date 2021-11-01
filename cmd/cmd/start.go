package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"wiki-link/config"
	"wiki-link/core/log"
	"wiki-link/initializers"
	"wiki-link/router"
	service "wiki-link/service/error"
)

var (
	configPath   string
	logPath      string
	mysqlLogPath string
	port         int
	StartCmd     = &cobra.Command{
		Use:     "run",
		Short:   "start run server",
		Example: "server run -port 8888 -config /etc/config/config.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

//初始化
func init() {
	StartCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config/config.yml", "config file path")
	StartCmd.PersistentFlags().StringVarP(&logPath, "log", "l", "./", "log path")
	StartCmd.PersistentFlags().StringVarP(&mysqlLogPath, "mysqllog", "m", "./mysql", "mysql log path")
}

//显示终端信息
func usage() {
	usageStr := `
	-------------------------------------------------------------------------------------------------
					*        *******   *        *   *  *
					*           *      *  *     *   * *
					*           *      *    *   *   *
					*           *      *      * *   * *
					*******  ********  *        *   *   *

				   welcome to use wikilink command
					    copyright @wikilink
	-------------------------------------------------------------------------------------------------
	`
	fmt.Printf("%s\n", usageStr)
}

//初始化各种服务
func setup() {
	if configPath != "" {
		config.ConfigFilePath = configPath
	}
	initializers.InitAllResources()
	//获取错误信息
	service.InitErrors()
}

//运行
func run() error {
	defer initializers.CloseResources()
	//初始化日志
	zapLog := log.Init("server", logPath)
	//获取路由
	engine := router.RouterEngine(zapLog)
	//启动
	engine.Run(fmt.Sprintf(":%v", port))
	return nil
}
