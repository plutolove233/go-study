package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{
		Topic: "delayTask",
		Value: sarama.StringEncoder("this is a test log"),
	}
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err=", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		println("send message failed, err=", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
