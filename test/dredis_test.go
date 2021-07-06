package test

import (
	"dragon/core/dragon"
	"dragon/core/dragon/dredis"
	"log"
	"testing"
)

func TestSet(t *testing.T) {
	dragon.AppInit()
	res, err := dredis.Redis.Set("async_task", 0, 0).Result()
	log.Println("res", res, "err", err)
}
