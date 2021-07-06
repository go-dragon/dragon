package rabbitmq

import (
	"github.com/streadway/amqp"
)

// Rabbit struct
type Rabbit struct {
	Dsn          string
	Conn         *amqp.Connection
	Channel      *amqp.Channel
	Queue        *amqp.Queue
	QueueName    string
	ExchangeName string
	RoutingKey   string
	PubConfirmCh chan amqp.Confirmation //publish confirm channel
}

// New a Rabbit
// dsn amqp://guest:guest@localhost:5672/
// exchangeKind: direct:精准推送；fanout:广播。推送到绑定到此交换机下的所有队列；topic组播。比如上面我绑定的关键字是sms_send，那么他可以推送到*.sms_send的所有队列。#表示匹配单个或者多个单词如：sms_send.#
func New(dsn string, exchangeName string, exchangeKind string, queueName string, routingKey string) (*Rabbit, error) {
	// connection
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	// channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// default publish msg with confirm
	err = ch.Confirm(false)
	if err != nil {
		return nil, err
	}
	pubConfirmCh := make(chan amqp.Confirmation)
	pubConfirmCh = ch.NotifyPublish(pubConfirmCh) // set publish notify

	// exchange
	err = ch.ExchangeDeclare(exchangeName, exchangeKind, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	// queue
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// bind queue
	err = ch.QueueBind(queueName, routingKey, exchangeName, false, nil)
	if err != nil {
		return nil, err
	}

	return &Rabbit{
		Dsn:          dsn,
		Conn:         conn,
		Channel:      ch,
		Queue:        &q,
		QueueName:    queueName,
		ExchangeName: exchangeName,
		RoutingKey:   routingKey,
		PubConfirmCh: pubConfirmCh,
	}, nil
}

// Close
func (r *Rabbit) Close() {
	r.Channel.Close()
	r.Conn.Close()
}

// Publish
func (r *Rabbit) Publish(body string) error {
	err := r.Channel.Publish(r.ExchangeName, r.RoutingKey, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         []byte(body),
	})
	return err
}

// Consumer, msg need to ack
func (r *Rabbit) GetConsumer(consumerName string) (<-chan amqp.Delivery, error) {
	return r.Channel.Consume(
		r.QueueName,  // queue
		consumerName, // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
}
