package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"time"
)

// produce content, if key = "" then kafka key is nil
func Produce(topic string, content string, key string) error {
	// to produce messages
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", viper.GetString("kafka.broker"), topic, partition)
	defer func() {
		conn.Close()
	}()
	conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	msg := kafka.Message{
		Value: []byte(content),
	}
	if key != "" {
		msg.Key = []byte(key)
	}
	_, err := conn.WriteMessages(
		msg,
	)
	return err
}

// to consume messages https://github.com/segmentio/kafka-go
func GetConsumerConn(topic string, offset int64) *kafka.Reader {
	// to consume messages
	// make a new reader that consumes from topic-A, partition 0
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{viper.GetString("kafka.broker")},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(offset)
	return r
}
