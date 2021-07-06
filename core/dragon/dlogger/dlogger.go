package dlogger

import (
	"bytes"
	"dragon/core/dragon/conf"
	"dragon/tools"
	"fmt"
	"github.com/go-dragon/util"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
	"time"
)

const (
	DebugLevel    = "DEBUG"
	InfoLevel     = "INFO"
	WarnLevel     = "WARN"
	ErrorLevel    = "FATAL_ERROR"
	SqlInfoLevel  = "SQL_INFO"
	SqlErrorLevel = "SQL_ERROR"
)

// 日志buffer，定时扫描刷到磁盘中
var logBuf = bytes.NewBufferString("")

// 日志缓存处理锁
var logBufMutex = sync.Mutex{}

// tick将日志写入文件中
func TickFlush() {
	tk := time.NewTicker(300 * time.Millisecond)
	defer tk.Stop()
	for range tk.C {
		// 取出缓存区日志，固化到本地
		logBufMutex.Lock()
		content := logBuf.String()
		bt, _ := tools.UnicodeToZh([]byte(content))
		logBuf.Reset()
		logBufMutex.Unlock()

		// 根据data类型删除json或者字符串
		now := time.Now()
		date := now.Format("2006-01-02")
		// 生成或打开文件
		logDir := viper.GetString("log.dir")
		path := conf.ExecDir + "/" + logDir + "/" + date + ".log"
		logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Println(fmt.Sprintf("error:%+v", err))
		}
		logFile.WriteString(string(bt))
		logFile.Close()
	}
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
	if data == nil || len(data) == 0 {
		// 如果data为空，不进行打印
		return
	}
	var logInfo string
	d, _ := util.FastJson.Marshal(&data)
	logInfo = string(d)
	// todo check if safe
	datetime := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("[%s] [%s] || %s \r\n\r\n", datetime, level, logInfo)
	// 写入缓冲区，日志
	logBufMutex.Lock()
	logBuf.WriteString(content)
	logBufMutex.Unlock()
}

func Debug(data ...interface{}) {
	writeLog(DebugLevel, data...)
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
