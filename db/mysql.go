package db

import (
	xormlogger "wiki-link/core/log/xorm"

	"github.com/arthurkiller/rollingwriter"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	engine *xorm.Engine
)

// Init 初始化
func InitMysqlConnect(driverName, mysqlDSN string, maxIdleConns, maxOpenConns int, logPath string) (err error) {
	engine, err = xorm.NewEngine(driverName, mysqlDSN)
	if err != nil {
		return err
	}

	if err := engine.Ping(); err != nil {
		return err
	}

	engine.SetMaxIdleConns(maxIdleConns)
	engine.SetMaxOpenConns(maxOpenConns)

	config := rollingwriter.Config{
		LogPath:                logPath,                     //日志路径
		TimeTagFormat:          "060102150405",              //时间格式串
		FileName:               "mysql",                     //日志文件名
		MaxRemain:              5,                           //配置日志最大存留数
		RollingPolicy:          rollingwriter.VolumeRolling, //配置滚动策略 norolling timerolling volumerolling
		RollingTimePattern:     "* * * * * *",               //配置时间滚动策略
		RollingVolumeSize:      "1M",                        //配置截断文件下限大小
		WriterMode:             "lock",
		BufferWriterThershould: 8 * 1024 * 1024,
		Compress:               true,
	}

	writer, err := rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		return err
	}

	var logger *xormlogger.SimpleLogger = xormlogger.NewSimpleLogger(writer)

	engine.SetLogger(logger)
	//engine.ShowSQL(true)

	return
}
