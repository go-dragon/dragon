package test

import (
	"dragon/core/dragon"
	"dragon/core/dragon/dlogger"
	"testing"
	"time"
)

func Test_writeLog(t *testing.T) {
	dragon.AppInit()
	//dlogger.Info("1", "2", "3")
	dlogger.Info(map[string]interface{}{
		"hello": "world",
		"x":     1,
	}, "666")
	time.Sleep(time.Second)
}
