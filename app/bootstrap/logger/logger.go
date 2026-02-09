package logger

import (
	"io"
	"os"
	"path/filepath"

	"ilicense-lite/config"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	appLogFileName     = "app.log"
	serviceLogFileName = "service.log"
)

// AppLogger 定义全局日志变量
var AppLogger *logrus.Logger
var ServiceLogger *logrus.Logger

// InitLogger 初始化日志
func InitLogger() {
	// 设置日志级别
	level, err := logrus.ParseLevel(config.Config.Log.Level)
	if err != nil {
		logrus.Fatalf("parse log level error: %v", err)
	}
	// 设置应用日志记录器
	AppLogger = logrus.New()
	AppLogger.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   filepath.Join(config.Config.Log.Dir, appLogFileName),
		MaxSize:    10,   // MB
		MaxBackups: 3,    // 保留旧文件的最大个数
		MaxAge:     1,    // 天
		Compress:   true, // 是否压缩/归档旧文件
	}))
	AppLogger.SetLevel(level)
	AppLogger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: false})
	AppLogger.SetReportCaller(true)
	// 设置业务日志记录器
	ServiceLogger = logrus.New()
	ServiceLogger.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   filepath.Join(config.Config.Log.Dir, serviceLogFileName),
		MaxSize:    10,   // MB
		MaxBackups: 3,    // 保留旧文件的最大个数
		MaxAge:     1,    // 天
		Compress:   true, // 是否压缩/归档旧文件
	}))
	ServiceLogger.SetLevel(level)
	ServiceLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})
	ServiceLogger.SetReportCaller(true)
}
