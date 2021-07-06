package test

import (
	"context"
	"dragon/core/dragon/conf"
	"dragon/tools/kafka"
	"fmt"
	"log"
	"testing"
)

// test produce
func TestProduce(t *testing.T) {
	conf.InitConf()
	err := kafka.Produce("test", "hello kafka~", "hello key")
	log.Println("kafka.Produce err", err)
}

// test Consumer
func TestConsume(t *testing.T) {
	// todo 本地业务处理时要注意固化记录offset,下次启动从offset开始
	conf.InitConf()
	var offset int64 = 25
	r := kafka.GetConsumerConn("test", offset)
	defer func() {
		r.Close()
	}()
	for {
		m, err := r.ReadMessage(context.Background())
		log.Println(22222)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		offset = m.Offset
	}
}
