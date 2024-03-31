package utils

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
)

type RocketMq struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var p rocketmq.Producer

// 初始化生产消息
func InitProducer(r *RocketMq) {
	p, _ = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{fmt.Sprintf("%s:%d", r.Host, r.Port)})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
}

// 生产消息
func Producer(topic string, message []byte) error {
	msg := primitive.NewMessage(topic, message)
	//1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	msg.WithDelayTimeLevel(3)
	_, err := p.SendSync(context.Background(), msg)
	return err
}

// 初始化消费消息
func InitConsumer(topic, group string, r *RocketMq) {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{fmt.Sprintf("%s:%d", r.Host, r.Port)})),
	)
	err := c.Subscribe(topic, consumer.MessageSelector{}, Consumer)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

// 消费者逻辑
func Consumer(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for i := range msgs {
		fmt.Printf("subscribe callback: %v \n", msgs[i])
	}
	return consumer.ConsumeSuccess, nil
}
