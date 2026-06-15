package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// 消费者，接收消息
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("创建消费者失败", err)
	}
	defer consumer.Close()

	// 订阅消息
	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("订阅分区失败", err)
	}
	defer partitionConsumer.Close()

	fmt.Println("消费者已启动，等待接收消息...")

	// 连接kafka，创建生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("创建生产者失败", err)
	}
	defer producer.Close()

	// 发送消息
	msg := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("hello kafka"),
	}

	Partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("发送消息失败", err)
	}

	fmt.Printf("消息发送成功！分区：%d，偏移量：%d\n", Partition, offset)

	// 等待接收消息
	select {
	case msg := <-partitionConsumer.Messages():
		fmt.Printf("接收到消息：%s\n", msg.Value)
	case err := <-partitionConsumer.Errors():
		log.Println("接收消息出错：", err)
	case <-time.After(time.Second * 5):
		fmt.Println("5秒后未收到消息")
	}
}
