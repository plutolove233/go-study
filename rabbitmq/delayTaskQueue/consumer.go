// 消费者.go
package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		println(err.Error())
		return
	}
	defer conn.Close()

	c, err := conn.Channel()
	if err != nil {
		println(err.Error())
		return
	}
	defer c.Close()

	err = c.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		println(err.Error())
		return
	}

	q, err := c.QueueDeclare(
		"test-logs",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	_, errDelay := c.QueueDeclare(
		"delayTask-test",
		false,
		false,
		true,
		false,
		amqp.Table{
			"x-dead-letter-exchange": "logs",
		})
	if errDelay != nil {
		panic(errDelay)
	}
	err = c.QueueBind(q.Name, "", "logs", false, nil)
	if err != nil {
		panic(err)
	}

	msgs, err := c.Consume(q.Name, "", true, false, false, false, nil) //监听dlxQueue队列
	if err != nil {
		println(err.Error())
		return
	}

	for d := range msgs {
		fmt.Printf("收到信息: %s\n", d.Body) // 收到消息，业务处理
	}
}
