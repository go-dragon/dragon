package dlogger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

const (
	InfoLevel     = "INFO"
	WarnLevel     = "WARN"
	ErrorLevel    = "FATAL_ERROR"
	SqlInfoLevel  = "SQL_INFO"
	SqlErrorLevel = "SQL_ERROR"
)

// LoggerZap
var LoggerZap, _ = zap.NewProduction()

// Enable enable to print logs
var Enable = false

func Init() {
	Enable = viper.GetBool("log.enable")
}

// checkFileIsExist
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// write log
func writeLog(level string, data ...interface{}) {
	if !Enable {
		return
	}
	if data == nil || len(data) == 0 {
		// 如果data为空，不进行打印
		return
	}
	defer LoggerZap.Sync()
	// 打印日志
	switch level {
	case SqlInfoLevel:
		LoggerZap.Info(SqlInfoLevel, zap.Any("body", data))
	case InfoLevel:
		LoggerZap.Info(InfoLevel, zap.Any("body", data))
	case WarnLevel:
		LoggerZap.Warn(WarnLevel, zap.Any("body", data))
	case SqlErrorLevel:
		LoggerZap.Error(SqlErrorLevel, zap.Any("body", data))
	case ErrorLevel:
		LoggerZap.Error(ErrorLevel, zap.Any("body", data))
	}
}

func Info(data ...interface{}) {
	writeLog(InfoLevel, data...)
}

func Warn(data ...interface{}) {
	writeLog(WarnLevel, data...)
}

func Error(data ...interface{}) {
	writeLog(ErrorLevel, data...)
}

func SqlInfo(data ...interface{}) {
	writeLog(SqlInfoLevel, data...)
}

func SqlError(data ...interface{}) {
	writeLog(SqlErrorLevel, data...)
}
