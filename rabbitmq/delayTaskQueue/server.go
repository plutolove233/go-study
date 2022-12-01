// 生产者.go
package main

import (
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	body := "hello world1"
	for i := 0; i < 10; i++ {
		err = ch.Publish(
			"",
			"delayTask-test",
			false,
			false,
			amqp.Publishing{
				Body:       []byte(body),
				Expiration: "5000", // 设置TTL为5秒
			})
		if err != nil {
			panic(err)
		}
	}
}
